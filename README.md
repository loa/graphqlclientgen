# graphqlclientgen

_Here Be Dragons_ :dragon:

This is an experiment of generating GraphQL clients in Golang. It's heavily inspired by [gqlgen](https://github.com/99designs/gqlgen).

- Goals
  - Structure config and generated clients in a future proof way to enable redesign of codegen
  - Single client generated for entire schema
  - ProtoClient interface to support multiple protocols and custom clients
  - Use static types in client interface to check api compability during build time
  - Enable field selection in client to partially fetch objects
- Limitations
  - No support for arguments on fields _(yet)_

## Get started

1. Add graphqlclientgen to `tools.go`
   ```golang
   //go:build tools

   package tools

   import (
     _ "github.com/loa/graphqlclientgen/cmd"
   )
   ```
2. Run go tidy to add graphqlclientgen as dependency
   ```bash
   go mod tidy
   ```
3. Create a new directory and run graphqlclientgen init
   ```bash
   mkdir exampleclient
   cd exampleclient

   go run github.com/loa/graphqlclientgen/cmd init --schema-path='../graph/*.graphqls'
   ```
4. Run generate _(init creates `exampleclient.go` with go gen comment)_
   ```bash
   go generate .
   ```

## Examples

```golang
func Example() {
  // create client, supports custom protocols through ProtoClient interface
  c := client.New(graphqlclientgen.NewHttpClient("http://localhost:8080/query"))

  todo, err := c.CreateTodo(context.TODO(),
    // static typed inputs
    client.NewTodo{
      Text:   "bar",
      UserId: "5",
    },
    // explicit requested fields, easy to use with auto-complete
    client.TodoFields{
      client.TodoFieldID,
      client.TodoFieldText,
      // supports requesting specific fields of related objects
      client.TodoFieldUser{
        client.UserFieldID,
        client.UserFieldName,
      },
    })
  }
```