module simple

go 1.25.0

require (
	github.com/99designs/gqlgen v0.17.91
	github.com/google/uuid v1.6.0
	github.com/loa/graphqlclientgen v0.17.0
	github.com/stretchr/testify v1.11.1
	github.com/vektah/gqlparser/v2 v2.5.35
)

require (
	github.com/agnivade/levenshtein v1.2.1 // indirect
	github.com/coder/websocket v1.8.15 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/go-viper/mapstructure/v2 v2.5.0 // indirect
	github.com/goccy/go-yaml v1.19.2 // indirect
	github.com/gorilla/websocket v1.5.3 // indirect
	github.com/hashicorp/golang-lru/v2 v2.0.7 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/sosodev/duration v1.4.0 // indirect
	github.com/urfave/cli/v3 v3.10.0 // indirect
	golang.org/x/exp v0.0.0-20260611194520-c48552f49976 // indirect
	golang.org/x/mod v0.37.0 // indirect
	golang.org/x/sync v0.21.0 // indirect
	golang.org/x/text v0.37.0 // indirect
	golang.org/x/tools v0.46.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/loa/graphqlclientgen => ./../../

tool (
	github.com/99designs/gqlgen
	github.com/loa/graphqlclientgen
)
