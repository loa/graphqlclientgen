package natsproto

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/loa/graphqlclientgen"
	"github.com/nats-io/nats.go"
)

type (
	Client struct {
		conn    *nats.Conn
		subject string
	}
)

// NewClient creates a new Client for graphqlclientgen clients
func NewClient(subject string, conn *nats.Conn) graphqlclientgen.ProtoClient {
	natsClient := Client{
		conn:    conn,
		subject: subject,
	}

	return natsClient
}

// Do performs nats request towards a GraphQL api
func (client Client) Do(ctx context.Context, in graphqlclientgen.Body, out any) error {
	b, err := json.Marshal(in)
	if err != nil {
		return err
	}

	res, err := client.conn.RequestWithContext(ctx, client.subject, b)
	if err != nil {
		return err
	}

	r := bytes.NewReader(res.Data)
	if err := json.NewDecoder(r).Decode(out); err != nil {
		return err
	}

	return nil
}
