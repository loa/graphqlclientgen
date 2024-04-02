package main_test

import (
	"context"
	"net/http/httptest"
	"testing"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/loa/graphqlclientgen"
	"github.com/loa/graphqlclientgen/examples/simple/client"
	"github.com/loa/graphqlclientgen/examples/simple/graph"
	"github.com/loa/graphqlclientgen/examples/simple/graph/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type SimpleSuite struct {
	suite.Suite

	client client.Client
	server *httptest.Server
}

func TestSimpleSuite(t *testing.T) {
	suite.Run(t, new(SimpleSuite))
}

func (suite *SimpleSuite) SetupTest() {
	gsrv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		DB: []*model.Todo{
			{ID: "0", Text: "foo"},
		},
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

func (suite *SimpleSuite) TestSimpleGetTodos() {
	todos, err := suite.client.Todos(context.TODO(),
		client.TodoFields{
			client.TodoFieldID,
			client.TodoFieldText,
		})
	assert.Nil(suite.T(), err)
	if assert.Nil(suite.T(), err) {
		if assert.Equal(suite.T(), 1, len(todos), "should exist 1 todo") {
			assert.Equal(suite.T(), todos[0].Text, "foo", "todo should be same")
		}
	}
}

func (suite *SimpleSuite) TestSimpleCreateTodos() {
	todo, err := suite.client.CreateTodo(context.TODO(), client.NewTodo{
		Text:   "bar",
		UserId: "5",
	}, client.TodoFields{
		client.TodoFieldID,
		client.TodoFieldText,
	})
	if assert.Nil(suite.T(), err) {
		assert.Equal(suite.T(), todo.Text, "bar", "todo should be same")
	}

	todos, err := suite.client.Todos(context.TODO(), client.TodoFields{
		client.TodoFieldID,
		client.TodoFieldText,
	})
	if assert.Nil(suite.T(), err) {
		assert.Equal(suite.T(), 2, len(todos), "should exist 2 todos")
	}
}
