package graphqlclientgen

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

func (client HttpClient) Do(body Body, in any) error {
	return nil
}
