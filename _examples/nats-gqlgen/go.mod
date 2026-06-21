module nats-gqlgen

go 1.25.0

require (
	github.com/99designs/gqlgen v0.17.91
	github.com/google/uuid v1.6.0
	github.com/loa/graphqlclientgen v0.18.0
	github.com/nats-io/nats-server/v2 v2.10.22
	github.com/nats-io/nats.go v1.52.0
	github.com/oklog/run v1.2.0
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
	github.com/klauspost/compress v1.18.6 // indirect
	github.com/minio/highwayhash v1.0.3 // indirect
	github.com/nats-io/jwt/v2 v2.5.8 // indirect
	github.com/nats-io/nkeys v0.4.16 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/sosodev/duration v1.4.0 // indirect
	github.com/urfave/cli/v3 v3.10.0 // indirect
	golang.org/x/crypto v0.53.0 // indirect
	golang.org/x/exp v0.0.0-20260611194520-c48552f49976 // indirect
	golang.org/x/mod v0.37.0 // indirect
	golang.org/x/sync v0.21.0 // indirect
	golang.org/x/sys v0.46.0 // indirect
	golang.org/x/text v0.38.0 // indirect
	golang.org/x/time v0.7.0 // indirect
	golang.org/x/tools v0.46.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/loa/graphqlclientgen => ./../../

tool (
	github.com/99designs/gqlgen
	github.com/loa/graphqlclientgen
)
