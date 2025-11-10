package graphqlclient

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
)

type (
	HttpClient struct {
		url    string
		client *http.Client
	}

	HttpClientOpts func(*HttpClient)
)

var _ ProtoClient = HttpClient{}

// NewHttpClient creates a new HttpClient for graphqlclientgen clients
func NewHttpClient(url string, opts ...HttpClientOpts) HttpClient {
	httpClient := HttpClient{
		url: url,
	}

	for _, opt := range opts {
		opt(&httpClient)
	}

	// set a default http.Client if not provided through options
	if httpClient.client == nil {
		httpClient.client = &http.Client{}
	}

	return httpClient
}

// WithHttpClient sets a custom http.Client for HttpClient
func WithHttpClient(c *http.Client) HttpClientOpts {
	return func(httpClient *HttpClient) {
		httpClient.client = c
	}
}

// Do performs http post request towards a GraphQL api
func (httpClient HttpClient) Do(ctx context.Context, in Body, out any) error {
	b, err := json.Marshal(in)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, httpClient.url, bytes.NewBuffer(b))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := httpClient.client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(out); err != nil {
		return err
	}

	return nil
}
