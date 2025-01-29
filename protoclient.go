package graphqlclientgen

import (
	"context"
	"encoding/json"
)

type (
	ProtoClient interface {
		Do(ctx context.Context, in Body, out any) error
	}

	Body struct {
		Query     string         `json:"query"`
		Variables map[string]any `json:"variables"`
	}

	Response struct {
		Data   json.RawMessage `json:"data"`
		Errors *Error          `json:"errors"`
	}
)
