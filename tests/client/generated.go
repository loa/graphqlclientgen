// Code generated by github.com/loa/graphqlclientgen, DO NOT EDIT.

package client

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/loa/graphqlclientgen"
)

type (
	// Client for graphqlclient
	Client struct {
		protoClient graphqlclientgen.ProtoClient
	}

	// Output
	Output struct {
		// Input
		Input string `json:"input"`
	}

	// OutputNillable
	OutputNillable struct {
		// Input
		Input *string `json:"input"`
	}

	OutputFieldScalar string
	OutputField       interface {
		OutputFieldGraphQL() string
	}
	OutputFields []OutputField

	OutputNillableFieldScalar string
	OutputNillableField       interface {
		OutputNillableFieldGraphQL() string
	}
	OutputNillableFields []OutputNillableField
)

var (
	OutputFieldInput         OutputFieldScalar         = "input"
	OutputNillableFieldInput OutputNillableFieldScalar = "input"
)

func (field OutputFieldScalar) OutputFieldGraphQL() string {
	return string(field)
}

func (field OutputNillableFieldScalar) OutputNillableFieldGraphQL() string {
	return string(field)
}

// New create new graphqlclient
func New(protoClient graphqlclientgen.ProtoClient) Client {
	return Client{
		protoClient: protoClient,
	}
}

// CustomError (query)
func (client Client) CustomError(
	ctx context.Context,
) (*bool, error) {
	fieldsContent := ""

	body := graphqlclientgen.Body{
		Query: fmt.Sprintf(`
        query {
            customError%s
        }`, fieldsContent),
		Variables: map[string]any{},
	}

	var res graphqlclientgen.Response
	if err := client.protoClient.Do(ctx, body, &res); err != nil {
		return nil, err
	}

	var data struct {
		CustomError *bool `json:"customError"`
	}

	if res.Errors != nil {
		return nil, *res.Errors
	}

	if err := json.Unmarshal(res.Data, &data); err != nil {
		return nil, err
	}

	return data.CustomError, nil
}

// ReturnScalar (query)
func (client Client) ReturnScalar(
	ctx context.Context,
	input bool,
) (bool, error) {
	fieldsContent := ""

	body := graphqlclientgen.Body{
		Query: fmt.Sprintf(`
        query ($input: Boolean!) {
            returnScalar (input: $input)%s
        }`, fieldsContent),
		Variables: map[string]any{
			"input": input,
		},
	}

	var res graphqlclientgen.Response
	if err := client.protoClient.Do(ctx, body, &res); err != nil {
		var placeholder bool
		return placeholder, err
	}

	var data struct {
		ReturnScalar bool `json:"returnScalar"`
	}

	if res.Errors != nil {
		return data.ReturnScalar, *res.Errors
	}

	if err := json.Unmarshal(res.Data, &data); err != nil {
		return data.ReturnScalar, err
	}

	return data.ReturnScalar, nil
}

// ReturnScalarNillable (query)
func (client Client) ReturnScalarNillable(
	ctx context.Context,
	input *bool,
) (*bool, error) {
	fieldsContent := ""

	body := graphqlclientgen.Body{
		Query: fmt.Sprintf(`
        query ($input: Boolean) {
            returnScalarNillable (input: $input)%s
        }`, fieldsContent),
		Variables: map[string]any{
			"input": input,
		},
	}

	var res graphqlclientgen.Response
	if err := client.protoClient.Do(ctx, body, &res); err != nil {
		return nil, err
	}

	var data struct {
		ReturnScalarNillable *bool `json:"returnScalarNillable"`
	}

	if res.Errors != nil {
		return nil, *res.Errors
	}

	if err := json.Unmarshal(res.Data, &data); err != nil {
		return nil, err
	}

	return data.ReturnScalarNillable, nil
}

// Simple (query)
func (client Client) Simple(
	ctx context.Context,
	fields OutputFields,
) (Output, error) {
	var s []string
	for _, field := range fields {
		s = append(s, field.OutputFieldGraphQL())
	}
	fieldsContent := fmt.Sprintf(" { %s }", strings.Join(s, ","))

	body := graphqlclientgen.Body{
		Query: fmt.Sprintf(`
        query {
            simple%s
        }`, fieldsContent),
		Variables: map[string]any{},
	}

	var res graphqlclientgen.Response
	if err := client.protoClient.Do(ctx, body, &res); err != nil {
		return Output{}, err
	}

	var data struct {
		Simple Output `json:"simple"`
	}

	if res.Errors != nil {
		return data.Simple, *res.Errors
	}

	if err := json.Unmarshal(res.Data, &data); err != nil {
		return data.Simple, err
	}

	return data.Simple, nil
}

// SimpleArgument (query)
func (client Client) SimpleArgument(
	ctx context.Context,
	input string,
	fields OutputFields,
) (Output, error) {
	var s []string
	for _, field := range fields {
		s = append(s, field.OutputFieldGraphQL())
	}
	fieldsContent := fmt.Sprintf(" { %s }", strings.Join(s, ","))

	body := graphqlclientgen.Body{
		Query: fmt.Sprintf(`
        query ($input: String!) {
            simpleArgument (input: $input)%s
        }`, fieldsContent),
		Variables: map[string]any{
			"input": input,
		},
	}

	var res graphqlclientgen.Response
	if err := client.protoClient.Do(ctx, body, &res); err != nil {
		return Output{}, err
	}

	var data struct {
		SimpleArgument Output `json:"simpleArgument"`
	}

	if res.Errors != nil {
		return data.SimpleArgument, *res.Errors
	}

	if err := json.Unmarshal(res.Data, &data); err != nil {
		return data.SimpleArgument, err
	}

	return data.SimpleArgument, nil
}

// SimpleArgumentNillable (query)
func (client Client) SimpleArgumentNillable(
	ctx context.Context,
	input *string,
	fields OutputNillableFields,
) (OutputNillable, error) {
	var s []string
	for _, field := range fields {
		s = append(s, field.OutputNillableFieldGraphQL())
	}
	fieldsContent := fmt.Sprintf(" { %s }", strings.Join(s, ","))

	body := graphqlclientgen.Body{
		Query: fmt.Sprintf(`
        query ($input: String) {
            simpleArgumentNillable (input: $input)%s
        }`, fieldsContent),
		Variables: map[string]any{
			"input": input,
		},
	}

	var res graphqlclientgen.Response
	if err := client.protoClient.Do(ctx, body, &res); err != nil {
		return OutputNillable{}, err
	}

	var data struct {
		SimpleArgumentNillable OutputNillable `json:"simpleArgumentNillable"`
	}

	if res.Errors != nil {
		return data.SimpleArgumentNillable, *res.Errors
	}

	if err := json.Unmarshal(res.Data, &data); err != nil {
		return data.SimpleArgumentNillable, err
	}

	return data.SimpleArgumentNillable, nil
}

// SimpleNillable (query)
func (client Client) SimpleNillable(
	ctx context.Context,
	fields OutputFields,
) (*Output, error) {
	var s []string
	for _, field := range fields {
		s = append(s, field.OutputFieldGraphQL())
	}
	fieldsContent := fmt.Sprintf(" { %s }", strings.Join(s, ","))

	body := graphqlclientgen.Body{
		Query: fmt.Sprintf(`
        query {
            simpleNillable%s
        }`, fieldsContent),
		Variables: map[string]any{},
	}

	var res graphqlclientgen.Response
	if err := client.protoClient.Do(ctx, body, &res); err != nil {
		return nil, err
	}

	var data struct {
		SimpleNillable *Output `json:"simpleNillable"`
	}

	if res.Errors != nil {
		return nil, *res.Errors
	}

	if err := json.Unmarshal(res.Data, &data); err != nil {
		return nil, err
	}

	return data.SimpleNillable, nil
}
