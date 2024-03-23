package graphqlclientgen

import "context"

type (
	HttpClient struct {
		URL string
	}
)

var _ ProtoClient = HttpClient{}

func NewHttpClient(url string) HttpClient {
	return HttpClient{
		URL: url,
	}
}

func (client HttpClient) Do(ctx context.Context, body Body, in any) error {
	return nil
}
