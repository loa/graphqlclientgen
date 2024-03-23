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

func (suite *SimpleSuite) SetupTest() {
	gsrv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		DB: []*model.Todo{
			{ID: "0", Text: "foobar"},
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

func (suite *SimpleSuite) TestSimpleGetTodos(t *testing.T) {
	todos, err := suite.client.Todos(context.TODO())

	if assert.Nil(t, err) {
		if assert.Equal(t, 1, len(todos), "should exist 1 todo") {
			assert.Equal(t, todos[0].Text, "foobar", "todo should be same")
		}
	}
}
