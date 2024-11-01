package main_test

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math/rand/v2"
	"testing"
	"time"

	"nats-gqlgen/client"
	"nats-gqlgen/graph"
	"nats-gqlgen/graph/model"

	"github.com/google/uuid"
	"github.com/loa/graphqlclientgen/natsproto"
	"github.com/nats-io/nats-server/v2/server"
	"github.com/nats-io/nats.go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type NatsGqlGenSuite struct {
	suite.Suite

	client client.Client
	server *server.Server

	db  *graph.DB
	sub *nats.Subscription
}

func TestNatsGqlGenSuite(t *testing.T) {
	suite.Run(t, new(NatsGqlGenSuite))
}

func (suite *NatsGqlGenSuite) SetupSuite() {
	srv, err := startNatsServer()
	if err != nil {
		log.Fatal(err)
	}

	suite.server = srv

	conn, err := nats.Connect(suite.server.ClientURL())
	if err != nil {
		log.Fatal(err)
	}
	suite.client = client.New(natsproto.NewClient(
		"example.query", conn,
	))
}

func (suite *NatsGqlGenSuite) TearDownSuite() {
	suite.server.Shutdown()
	suite.server.WaitForShutdown()
}

func (suite *NatsGqlGenSuite) SetupTest() {
	suite.db = &graph.DB{
		Todos: []*model.Todo{},
		Users: []*model.User{},
	}

	for range rand.IntN(2) + 1 {
		id := uuid.Must(uuid.NewV7())
		suite.db.Users = append(suite.db.Users, &model.User{
			ID:   id,
			Name: fmt.Sprintf("User #%s", id.String()),
		})
	}

	for _, user := range suite.db.Users {
		for range rand.IntN(10) + 1 {
			suite.db.TodosIncrementalID += 1
			suite.db.Todos = append(suite.db.Todos, &model.Todo{
				ID:     fmt.Sprint(suite.db.TodosIncrementalID),
				Text:   "Some static text",
				UserID: user.ID,
			})
		}
	}

	conn, err := nats.Connect(suite.server.ClientURL())
	if err != nil {
		log.Fatal(err)
	}

	gsrv := natsproto.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		DB: suite.db,
	}}))

	suite.sub, err = conn.Subscribe("example.query", gsrv.HandleFunc)
	if err != nil {
		log.Fatal(err)
	}
}

func (suite *NatsGqlGenSuite) TearDownTest() {
	suite.sub.Drain()
}

func (suite *NatsGqlGenSuite) TestGetTodoByID() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	expected := suite.db.Todos[rand.IntN(suite.db.TodosIncrementalID)]

	todo, err := suite.client.Todo(ctx, expected.ID, client.TodoFields{
		client.TodoFieldID,
		client.TodoFieldText,
	})

	if assert.Nil(suite.T(), err) {
		assert.Equal(suite.T(), expected.ID, todo.ID, "should be equal")
		assert.Equal(suite.T(), expected.Text, todo.Text, "should be equal")
	}
}

func (suite *NatsGqlGenSuite) TestGetTodoByIDNotFound() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err := suite.client.Todo(ctx, fmt.Sprint(suite.db.TodosIncrementalID+100), client.TodoFields{
		client.TodoFieldID,
		client.TodoFieldText,
	})

	assert.EqualError(suite.T(), err, "not found")
}

func (suite *NatsGqlGenSuite) TestGetTodos() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	todos, err := suite.client.Todos(ctx,
		client.TodoFields{
			client.TodoFieldID,
			client.TodoFieldText,
		})

	if assert.Nil(suite.T(), err) {
		assert.Equal(suite.T(), len(suite.db.Todos), len(todos), "should be equal")
	}
}

func (suite *NatsGqlGenSuite) TestCreateTodos() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	expectedLen := len(suite.db.Todos) + 1

	todo, err := suite.client.CreateTodo(ctx,
		client.NewTodo{
			Text: "New todo!",
			User: suite.db.Users[0].ID,
		},
		client.TodoFields{
			client.TodoFieldID,
			client.TodoFieldText,
		})

	if assert.Nil(suite.T(), err) {
		assert.Equal(suite.T(), todo.Text, "New todo!", "should be equal")
	}

	todos, err := suite.client.Todos(ctx,
		client.TodoFields{
			client.TodoFieldID,
			client.TodoFieldText,
			client.UserFields{
				client.UserFieldID,
				client.UserFieldName,
			},
		})

	if assert.Nil(suite.T(), err) {
		assert.Equal(suite.T(), expectedLen, len(todos), "should be equal")
	}
}

func startNatsServer() (*server.Server, error) {
	opts := &server.Options{
		NoSigs: true,
	}

	srv, err := server.NewServer(opts)
	if err != nil {
		return nil, err
	}

	srv.Start()
	if !srv.ReadyForConnections(5 * time.Second) {
		return nil, errors.New("nats server start timeout")
	}

	return srv, nil
}
