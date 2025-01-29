package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/loa/graphqlclientgen"
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

func TestNatsReturnScalar(t *testing.T) {
	c, teardown := setupGraph(t)
	defer teardown()

	tests := []bool{true, false}

	for _, test := range tests {
		t.Run(fmt.Sprint(test), func(t *testing.T) {
			actual, err := c.ReturnScalar(context.TODO(), test)

			require.Nil(t, err)
			require.Equal(t, test, actual)
		})
	}
}

func TestNatsReturnScalarNillable(t *testing.T) {
	c, teardown := setupGraph(t)
	defer teardown()

	test1, test2 := true, false
	tests := []*bool{&test1, &test2, nil}

	for _, test := range tests {
		t.Run(boolPointerValue(test), func(t *testing.T) {
			actual, err := c.ReturnScalarNillable(context.TODO(), test)

			require.Nil(t, err)
			require.Equal(t, test, actual)
		})
	}
}

func TestNatsCustomError(t *testing.T) {
	c, teardown := setupGraph(t)
	defer teardown()

	_, err := c.CustomError(context.TODO())

	var gerr graphqlclientgen.Error
	require.ErrorAs(t, err, &gerr)
	require.True(t, gerr.ExtensionEqualString("code", "NOT_FOUND"))
}
