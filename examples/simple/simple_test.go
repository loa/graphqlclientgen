package main_test

import (
	"context"
	"fmt"
	"math/rand/v2"
	"net/http/httptest"
	"testing"

	"simple/client"
	"simple/graph"
	"simple/graph/model"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/google/uuid"
	"github.com/loa/graphqlclientgen"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type SimpleSuite struct {
	suite.Suite

	db     *graph.DB
	client client.Client
	server *httptest.Server
}

func TestSimpleSuite(t *testing.T) {
	suite.Run(t, new(SimpleSuite))
}

func (suite *SimpleSuite) SetupTest() {
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

	gsrv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		DB: suite.db,
	}}))

	suite.server = httptest.NewServer(gsrv)

	suite.client = client.New(graphqlclientgen.NewHttpClient(
		suite.server.URL,
		graphqlclientgen.WithHttpClient(suite.server.Client()),
	))
}

func (suite *SimpleSuite) TearDownTest() {
	suite.server.Close()
}

func (suite *SimpleSuite) TestSimpleGetTodoByID() {
	expected := suite.db.Todos[rand.IntN(suite.db.TodosIncrementalID)]

	todo, err := suite.client.Todo(context.TODO(), expected.ID, client.TodoFields{
		client.TodoFieldID,
		client.TodoFieldText,
	})

	if assert.Nil(suite.T(), err) {
		assert.Equal(suite.T(), expected.ID, todo.ID, "should be equal")
		assert.Equal(suite.T(), expected.Text, todo.Text, "should be equal")
	}
}

func (suite *SimpleSuite) TestSimpleGetTodos() {
	todos, err := suite.client.Todos(context.TODO(),
		client.TodoFields{
			client.TodoFieldID,
			client.TodoFieldText,
		})

	if assert.Nil(suite.T(), err) {
		assert.Equal(suite.T(), len(suite.db.Todos), len(todos), "should be equal")
	}
}

func (suite *SimpleSuite) TestSimpleCreateTodos() {
	expectedLen := len(suite.db.Todos) + 1

	todo, err := suite.client.CreateTodo(context.TODO(),
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

	todos, err := suite.client.Todos(context.TODO(), client.TodoFields{
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
