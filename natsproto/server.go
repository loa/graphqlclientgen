package natsproto

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/nats-io/nats.go"
	"github.com/vektah/gqlparser/v2/ast"
	"github.com/vektah/gqlparser/v2/gqlerror"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/executor"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
)

type (
	Server struct {
		exec *executor.Executor
	}
)

// NewServer creates a new Server for nats subscription
func NewServer(es graphql.ExecutableSchema) *Server {
	return &Server{
		exec: executor.New(es),
	}
}

// NewDefaultServer creates a new Server with sensible defaults for nats subscription
func NewDefaultServer(es graphql.ExecutableSchema) *Server {
	server := NewServer(es)

	server.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	server.Use(extension.Introspection{})
	server.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	return server
}

func (server *Server) AddTransport(transport graphql.Transport) {}

func (server *Server) SetErrorPresenter(f graphql.ErrorPresenterFunc) {
	server.exec.SetErrorPresenter(f)
}

func (server *Server) SetRecoverFunc(f graphql.RecoverFunc) {
	server.exec.SetRecoverFunc(f)
}

func (server *Server) SetQueryCache(cache graphql.Cache[*ast.QueryDocument]) {
	server.exec.SetQueryCache(cache)
}

func (server *Server) SetParserTokenLimit(limit int) {
	server.exec.SetParserTokenLimit(limit)
}

func (server *Server) Use(extension graphql.HandlerExtension) {
	server.exec.Use(extension)
}

// AroundFields is a convenience method for creating an extension that only implements field middleware
func (server *Server) AroundFields(f graphql.FieldMiddleware) {
	server.exec.AroundFields(f)
}

// AroundRootFields is a convenience method for creating an extension that only implements field middleware
func (server *Server) AroundRootFields(f graphql.RootFieldMiddleware) {
	server.exec.AroundRootFields(f)
}

// AroundOperations is a convenience method for creating an extension that only implements operation middleware
func (server *Server) AroundOperations(f graphql.OperationMiddleware) {
	server.exec.AroundOperations(f)
}

// AroundResponses is a convenience method for creating an extension that only implements response middleware
func (server *Server) AroundResponses(f graphql.ResponseMiddleware) {
	server.exec.AroundResponses(f)
}

// HandleFunc handler for nats.Subscribe
func (server *Server) HandleFunc(msg *nats.Msg) {
	ctx := context.Background()
	defer func() {
		if err := recover(); err != nil {
			err := server.exec.PresentRecoveredError(ctx, err)
			gqlErr, _ := err.(*gqlerror.Error)
			resp := &graphql.Response{Errors: []*gqlerror.Error{gqlErr}}
			b, _ := json.Marshal(resp)
			msg.Respond(b)
		}
	}()

	ctx = graphql.StartOperationTrace(ctx)

	params := &graphql.RawParams{}
	start := graphql.Now()
	params.Headers = http.Header(msg.Header)
	params.ReadTime = graphql.TraceTiming{
		Start: start,
		End:   graphql.Now(),
	}

	dec := json.NewDecoder(bytes.NewReader(msg.Data))
	dec.UseNumber()

	if err := dec.Decode(&params); err != nil {
		gqlErr := gqlerror.Errorf(
			"msg json body could not be decoded: %+v body:%s",
			err,
			string(msg.Data),
		)
		resp := server.exec.DispatchError(ctx, gqlerror.List{gqlErr})
		writeJson(msg, resp)
		return
	}

	rc, opErr := server.exec.CreateOperationContext(ctx, params)
	if opErr != nil {
		resp := server.exec.DispatchError(graphql.WithOperationContext(ctx, rc), opErr)
		writeJson(msg, resp)
		return
	}

	var responses graphql.ResponseHandler
	responses, ctx = server.exec.DispatchOperation(ctx, rc)
	writeJson(msg, responses(ctx))
}

func writeJson(msg *nats.Msg, response *graphql.Response) {
	b, err := json.Marshal(response)
	if err != nil {
		writeJson(msg, &graphql.Response{Errors: gqlerror.List{{Message: err.Error()}}})
		return
	}

	err = msg.Respond(b)
	if errors.Is(err, nats.ErrMaxPayload) {
		writeJson(msg, &graphql.Response{Errors: gqlerror.List{{Message: err.Error()}}})
		return
	}

	// TODO: check for more nats errors that can be handled gracefully, on panic the requesting
	//			 client will hang and wait for a response that will never be received
	if err != nil {
		panic(err)
	}
}
