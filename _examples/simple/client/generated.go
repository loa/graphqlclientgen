// Code generated by github.com/loa/graphqlclientgen, DO NOT EDIT.

package client

import (
	"github.com/loa/graphqlclientgen"
)

type (
	// Client for graphqlclient
	Client struct {
		protoClient graphqlclientgen.ProtoClient
	}

	// Todo entry with text and done status
	Todo struct {
		// Done status of todo
		Done bool `json:"done"`
		// ID primary id of todo
		ID string `json:"id"`
		// Text todo text
		Text string `json:"text"`
		// User assigned to todo
		User *User `json:"user"`
	}

	// User with name and assigned todos
	User struct {
		// ID primary id of user
		ID string `json:"id"`
		// Name of user
		Name string `json:"name"`
		// Todos all todos assigned to user
		Todos *Todo `json:"todos"`
	}
)

// New create new graphqlclient
func New(protoClient graphqlclientgen.ProtoClient) Client {
	return Client{
		protoClient: protoClient,
	}
}

// CreateTodo (mutation) create a new todo
func (client Client) CreateTodo() (Todo, error) {
	body := graphqlclientgen.Body{
		Query: "mutation { createTodo {}}",
	}

	var res Todo

	if err := client.protoClient.Do(body, &res); err != nil {
		return Todo{}, err
	}

	return res, nil
}

// Todos (query) returns all todos
func (client Client) Todos() ([]Todo, error) {
	body := graphqlclientgen.Body{
		Query: "query { todos {}}",
	}

	var res []Todo

	if err := client.protoClient.Do(body, &res); err != nil {
		return nil, err
	}

	return res, nil
}
