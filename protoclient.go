package graphqlclientgen

type (
	ProtoClient interface {
		Do(body Body, in any) error
	}

	Body struct {
		Query     string `json:"query"`
		Variables string `json:"variables"`
	}
)
