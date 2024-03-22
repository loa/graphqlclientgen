package client

type (
	// Client for graphqlclient
	Client struct{}
)

// New creates new graphqlclient
func New() Client {
	return Client{}
}

type (
	Todo struct{}
)

// Todos query function
func (client Client) Todos() (Todo, error) {
	// query
	return Todo{}, nil
}

// CreateTodo mutation function
func (client Client) CreateTodo() (Todo, error) {
	// mutation
	return Todo{}, nil
}
