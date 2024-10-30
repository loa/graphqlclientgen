// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"github.com/google/uuid"
)

type Mutation struct {
}

type NewTodo struct {
	// todo text
	Text string `json:"text"`
	// user to assign todo
	User uuid.UUID `json:"user"`
}

type Query struct {
}

// Todo entry with text and done status
type Todo struct {
	// primary id of todo
	ID string `json:"id"`
	// todo text
	Text string `json:"text"`
	// done status of todo
	Done bool `json:"done"`
	// user assigned to todo
	User *User `json:"user"`
	// Todo belongs to a User
	UserID uuid.UUID `json:"-"`
}

// User with name and assigned todos
type User struct {
	// primary id of user
	ID uuid.UUID `json:"id"`
	// name of user
	Name string `json:"name"`
	// all todos assigned to user
	Todos []*Todo `json:"todos"`
}
