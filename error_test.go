package graphqlclientgen

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func ExampleError() {
	err := Error{
		{
			Message: "entry with incorrect type",
			Extensions: map[string]json.RawMessage{
				"featureA": json.RawMessage(`"true"`),
			},
		},
		{
			Message: "entry with correct type",
			Extensions: map[string]json.RawMessage{
				"featureA": json.RawMessage("true"),
			},
		},
	}

	fmt.Println(err.Error())
	// Output: entry with incorrect type
	// entry with correct type
}

func TestErrorUnwrap(t *testing.T) {
	err := Error{
		{
			Message: "entry with incorrect type",
			Extensions: map[string]json.RawMessage{
				"featureA": json.RawMessage(`"true"`),
			},
		},
		{
			Message: "entry with correct type",
			Extensions: map[string]json.RawMessage{
				"featureA": json.RawMessage("true"),
			},
		},
	}

	var expected []error
	for _, e := range err {
		expected = append(expected, e)
	}

	require.Equal(t, expected, err.Unwrap())
}

func TestErrorExtensionBool(t *testing.T) {
	input := Error{
		{
			Message: "entry with incorrect type",
			Extensions: map[string]json.RawMessage{
				"featureA": json.RawMessage(`"true"`),
			},
		},
		{
			Message: "entry with correct type",
			Extensions: map[string]json.RawMessage{
				"featureA": json.RawMessage("true"),
			},
		},
	}

	tests := []struct {
		Key string
		Val bool
		Ok  bool
	}{
		{Key: "featureA", Val: true, Ok: true},
		{Key: "dontexist", Val: false, Ok: false},
	}

	for _, test := range tests {
		t.Run(test.Key, func(t *testing.T) {
			var actual Error
			require.ErrorAs(t, input, &actual)

			val, ok := actual.ExtensionBool(test.Key)
			require.Equal(t, test.Val, val)
			require.Equal(t, test.Ok, ok)

			ok = actual.ExtensionEqualBool(test.Key, test.Val)
			require.Equal(t, test.Ok, ok)
		})
	}
}

func TestErrorExtensionFloat(t *testing.T) {
	input := Error{
		{
			Message: "entry with incorrect type",
			Extensions: map[string]json.RawMessage{
				"temp": json.RawMessage(`"true"`),
			},
		},
		{
			Message: "entry with correct type",
			Extensions: map[string]json.RawMessage{
				"temp": json.RawMessage("21.5"),
			},
		},
	}

	tests := []struct {
		Key string
		Val float64
		Ok  bool
	}{
		{Key: "temp", Val: 21.5, Ok: true},
		{Key: "dontexist", Val: 0.0, Ok: false},
	}

	for _, test := range tests {
		t.Run(test.Key, func(t *testing.T) {
			var actual Error
			require.ErrorAs(t, input, &actual)

			val, ok := actual.ExtensionFloat64(test.Key)
			require.InDelta(t, test.Val, val, 0.01)
			require.Equal(t, test.Ok, ok)
		})
	}
}

func TestErrorExtensionInt(t *testing.T) {
	input := Error{
		{
			Message: "entry with incorrect type",
			Extensions: map[string]json.RawMessage{
				"errorCode": json.RawMessage(`"not an int"`),
			},
		},
		{
			Message: "entry with incorrect type",
			Extensions: map[string]json.RawMessage{
				"errorCode": json.RawMessage(`3.14`),
			},
		},
		{
			Message: "entry with string value",
			Extensions: map[string]json.RawMessage{
				"errorCode": json.RawMessage(`404`),
			},
		},
	}

	tests := []struct {
		Key string
		Val int
		Ok  bool
	}{
		{Key: "errorCode", Val: 404, Ok: true},
		{Key: "dontexist", Val: 0, Ok: false},
	}

	for _, test := range tests {
		t.Run(test.Key, func(t *testing.T) {
			var actual Error
			require.ErrorAs(t, input, &actual)

			val, ok := actual.ExtensionInt(test.Key)
			require.Equal(t, test.Val, val)
			require.Equal(t, test.Ok, ok)

			ok = actual.ExtensionEqualInt(test.Key, test.Val)
			require.Equal(t, test.Ok, ok)
		})
	}
}

func TestErrorExtensionString(t *testing.T) {
	input := Error{
		{
			Message: "non string extension entry",
			Extensions: map[string]json.RawMessage{
				"code":   json.RawMessage(`1234`),
				"int":    json.RawMessage(`1234`),
				"number": json.RawMessage(`1234.5678`),
			},
		},
		{
			Message: "entry with string value",
			Extensions: map[string]json.RawMessage{
				"code": json.RawMessage(`"HELLO_WORLD"`),
			},
		},
	}

	tests := []struct {
		Key string
		Val string
		Ok  bool
	}{
		{Key: "code", Val: "HELLO_WORLD", Ok: true},
		{Key: "dontexist", Val: "", Ok: false},
		{Key: "int", Val: "", Ok: false},
		{Key: "number", Val: "", Ok: false},
	}

	for _, test := range tests {
		t.Run(test.Val, func(t *testing.T) {
			var actual Error
			require.ErrorAs(t, input, &actual)

			val, ok := actual.ExtensionString(test.Key)
			require.Equal(t, test.Val, val)
			require.Equal(t, test.Ok, ok)
		})
	}
}

func TestErrorExtensionEqualString(t *testing.T) {
	input := Error{
		{
			Message: "non string extension entry",
			Extensions: map[string]json.RawMessage{
				"code":   json.RawMessage(`1234`),
				"int":    json.RawMessage(`1234`),
				"number": json.RawMessage(`1234.5678`),
			},
		},
		{
			Message: "first entry with string value",
			Extensions: map[string]json.RawMessage{
				"code": json.RawMessage(`"HELLO_WORLD"`),
			},
		},
		{
			Message: "second entry with string value",
			Extensions: map[string]json.RawMessage{
				"code": json.RawMessage(`"FOO_BAR"`),
			},
		},
	}

	tests := []struct {
		Key string
		Val string
		Ok  bool
	}{
		{Key: "code", Val: "HELLO_WORLD", Ok: true},
		{Key: "code", Val: "FOO_BAR", Ok: true},
		{Key: "dontexist", Val: "", Ok: false},
		{Key: "int", Val: "1234", Ok: false},
		{Key: "number", Val: "1234.5678", Ok: false},
	}

	for _, test := range tests {
		t.Run(test.Val, func(t *testing.T) {
			var actual Error
			require.ErrorAs(t, input, &actual)

			ok := actual.ExtensionEqualString(test.Key, test.Val)
			require.Equal(t, test.Ok, ok)
		})
	}
}

func TestErrorExtensionUnmarshal(t *testing.T) {
	input := Error{
		{
			Extensions: map[string]json.RawMessage{
				"bool": json.RawMessage(`true`),
			},
		},
		{
			Extensions: map[string]json.RawMessage{
				"int": json.RawMessage(`42`),
				"map": json.RawMessage(`{"foo": "foo", "five": 5}`),
			},
		},
	}

	var actual Error
	require.ErrorAs(t, input, &actual)

	var i int
	require.ErrorIs(t, actual.ExtensionUnmarshal("dontexist", &i), ErrExtensionNotFound)

	require.Nil(t, actual.ExtensionUnmarshal("int", &i))
	require.Equal(t, 42, i)

	var b bool
	require.Nil(t, actual.ExtensionUnmarshal("bool", &b))
	require.Equal(t, true, b)

	expertedS := struct {
		Foo  string
		Five int
	}{
		Foo:  "foo",
		Five: 5,
	}
	actualS := struct {
		Foo  string
		Five int
	}{}
	require.Nil(t, actual.ExtensionUnmarshal("map", &actualS))
	require.Equal(t, expertedS, actualS)
}
