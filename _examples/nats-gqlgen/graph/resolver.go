package graph

import "nats-gqlgen/graph/model"

//go:generate go run github.com/99designs/gqlgen generate

type DB struct {
	Todos              []*model.Todo
	TodosIncrementalID int
	Users              []*model.User
}

type Resolver struct {
	DB *DB
}
