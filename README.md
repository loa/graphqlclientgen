# graphqlclientgen

_Here Be Dragons_ :dragon:

This is an experiment of generating GraphQL clients in Golang. It's heavily inspired by [gqlgen](https://github.com/99designs/gqlgen).

- Goals
  - Structure config and generated clients in a future proof way to enable redesign of codegen
  - Single client generated for entire schema
  - ProtoClient interface to support multiple protocols and custom clients
  - Use static types in client interface to check api compability during build time
  - Use GraphQL comments in Golang type and function definitions
  - Enable field selection in client to partially fetch objects
  - Reuse gqlgen custom scalars
- Limitations
  - Mostly designed to use GraphQL as standard RPC
  - Weak support for deep graph fetching _(can't have multiple fields with same object type)_
  - Only supports scalars and objects
  - Only supports hard-coded basic scalars _(bool, float, id, string, int)_

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
        client.UserFields{
          client.UserFieldID,
          client.UserFieldName,
        },
      })
    }
  ```