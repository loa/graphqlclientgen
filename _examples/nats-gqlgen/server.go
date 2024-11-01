package main

import (
	"context"
	"log"
	"log/slog"
	"nats-gqlgen/graph"
	"nats-gqlgen/graph/model"
	"net/http"
	"os"
	"syscall"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/loa/graphqlclientgen/natsproto"
	"github.com/nats-io/nats.go"
	"github.com/oklog/run"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	executableSchema := graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		DB: &graph.DB{
			Todos: []*model.Todo{},
			Users: []*model.User{},
		},
	}})

	var g run.Group
	{
		conn, err := nats.Connect(nats.DefaultURL)
		if err != nil {
			slog.Error("nats connect", "error", err)
			os.Exit(1)
		}

		gsrv := natsproto.NewDefaultServer(executableSchema)

		sub, err := conn.Subscribe("example.query", gsrv.HandleFunc)
		if err != nil {
			slog.Error("nats subscribe", "error", err)
			os.Exit(1)
		}

		done := make(chan struct{})
		g.Add(func() error {
			<-done
			sub.Drain()
			conn.Drain()
			return nil
		}, func(err error) {
			close(done)
		})
	}
	{
		httpServer := http.Server{Addr: ":" + port}

		srv := handler.NewDefaultServer(executableSchema)
		http.Handle("/", playground.Handler("GraphQL playground", "/query"))
		http.Handle("/query", srv)

		g.Add(
			func() error {
				log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
				return httpServer.ListenAndServe()
			},
			func(err error) {
				ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
				defer cancel()

				if err := httpServer.Shutdown(ctx); err != nil {
					slog.Error("error on http shutdown", "error", err)
				}
			},
		)
	}

	g.Add(run.SignalHandler(context.Background(), os.Interrupt, syscall.SIGTERM))
	if err := g.Run(); err != nil {
		log.Fatal(err)
	}
}
