# graphqlclientgen

_Here Be Dragons_

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
  - Weak support for deep graph fetching
  - Only supports scalars and objects
  - Only supports hard-coded basic scalars _(bool, float, id, string, int)_