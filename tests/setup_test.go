package tests

import (
	"testing"
	"time"

	"github.com/loa/graphqlclientgen/natsproto"
	"github.com/loa/graphqlclientgen/tests/client"
	"github.com/loa/graphqlclientgen/tests/graph"
	"github.com/nats-io/nats-server/v2/server"
	"github.com/nats-io/nats.go"
)

func setupGraph(t *testing.T) (client.Client, func()) {
	conn, cleanup := setupNatsServer(t)

	c := client.New(natsproto.NewClient("query", conn))
	gql := natsproto.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
		Resolvers: &graph.Resolver{},
	}))
	conn.Subscribe("query", gql.HandleFunc)

	return c, func() {
		cleanup()
	}
}

func setupNatsServer(t *testing.T) (*nats.Conn, func()) {
	opts := &server.Options{
		NoSigs: true,
	}

	srv, err := server.NewServer(opts)
	if err != nil {
		t.Fatal(err)
	}

	srv.Start()
	if !srv.ReadyForConnections(5 * time.Second) {
		t.Fatal("nats server start timeout")
	}

	conn, err := nats.Connect(srv.ClientURL())
	if err != nil {
		t.Fatal(err)
	}

	return conn, func() {
		conn.Close()
		srv.Shutdown()
	}
}
