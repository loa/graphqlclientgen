package graph

import "github.com/loa/graphqlclientgen/examples/simple/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	DB            []*model.Todo
	IncrementalID int
}
