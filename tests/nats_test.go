package tests

import (
	"context"
	"testing"

	"github.com/loa/graphqlclientgen/tests/client"
	"github.com/stretchr/testify/require"
)

func TestNatsSimple(t *testing.T) {
	c, teardown := setupGraph(t)
	defer teardown()

	simple, err := c.Simple(context.TODO(), client.OutputFields{
		client.OutputFieldInput,
	})

	require.Nil(t, err)
	require.Equal(t, "simple", simple.Input)
}

func TestNatsSimpleNillable(t *testing.T) {
	c, teardown := setupGraph(t)
	defer teardown()

	simple, err := c.SimpleNillable(context.TODO(), client.OutputFields{
		client.OutputFieldInput,
	})

	require.Nil(t, err)
	require.Equal(t, "simpleNillable", simple.Input)
}

func TestNatsSimpleArgument(t *testing.T) {
	c, teardown := setupGraph(t)
	defer teardown()

	tests := []string{"simple", "foo", "bar"}

	for _, test := range tests {
		t.Run(test, func(t *testing.T) {
			simple, err := c.SimpleArgument(context.TODO(), test, client.OutputFields{
				client.OutputFieldInput,
			})

			require.Nil(t, err)
			require.Equal(t, test, simple.Input)
		})
	}
}

func TestNatsSimpleArgumentNillable(t *testing.T) {
	c, teardown := setupGraph(t)
	defer teardown()

	simple, foo, bar := "simple", "foo", "bar"
	tests := []*string{&simple, &foo, &bar, nil}

	for _, test := range tests {
		t.Run(stringPointerValue(test), func(t *testing.T) {
			simple, err := c.SimpleArgumentNillable(context.TODO(), test, client.OutputNillableFields{
				client.OutputNillableFieldInput,
			})

			require.Nil(t, err)
			require.Equal(t, test, simple.Input)
		})
	}
}
