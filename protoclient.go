package graphqlclientgen

import "context"

type (
	ProtoClient interface {
		Do(ctx context.Context, body Body, in any) error
	}

	Body struct {
		Query     string         `json:"query"`
		Variables map[string]any `json:"variables"`
	}
)
