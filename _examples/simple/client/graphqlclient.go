package client

import (
	"github.com/loa/graphqlclientgen"
)

type (
	// Client for graphqlclient
	Client struct {
		protoClient graphqlclientgen.ProtoClient
	}
)

// New create new graphqlclient
func New(protoClient graphqlclientgen.ProtoClient) Client {
	return Client{
		protoClient: protoClient,
	}
}

type (
	Todo struct{}
)

// Todos query function
func (client Client) Todos() ([]Todo, error) {
	body := graphqlclientgen.Body{
		Query: "query {}",
	}

	var res []Todo

	if err := client.protoClient.Do(body, &res); err != nil {
		return nil, err
	}

	return res, nil
}

// CreateTodo mutation function
func (client Client) CreateTodo() (Todo, error) {
	body := graphqlclientgen.Body{
		Query: "mutation {}",
	}

	var res Todo

	if err := client.protoClient.Do(body, &res); err != nil {
		return Todo{}, err
	}

	return res, nil
}
