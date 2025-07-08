module nats-gqlgen

go 1.23.0

require (
	github.com/99designs/gqlgen v0.17.76
	github.com/google/uuid v1.6.0
	github.com/loa/graphqlclientgen v0.16.0
	github.com/nats-io/nats-server/v2 v2.10.22
	github.com/nats-io/nats.go v1.43.0
	github.com/oklog/run v1.2.0
	github.com/stretchr/testify v1.10.0
	github.com/vektah/gqlparser/v2 v2.5.30
)

require (
	github.com/agnivade/levenshtein v1.2.1 // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.7 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/go-viper/mapstructure/v2 v2.3.0 // indirect
	github.com/gorilla/websocket v1.5.3 // indirect
	github.com/hashicorp/golang-lru/v2 v2.0.7 // indirect
	github.com/klauspost/compress v1.18.0 // indirect
	github.com/minio/highwayhash v1.0.3 // indirect
	github.com/nats-io/jwt/v2 v2.5.8 // indirect
	github.com/nats-io/nkeys v0.4.11 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/sosodev/duration v1.3.1 // indirect
	github.com/urfave/cli/v2 v2.27.7 // indirect
	github.com/urfave/cli/v3 v3.3.8 // indirect
	github.com/xrash/smetrics v0.0.0-20250705151800-55b8f293f342 // indirect
	golang.org/x/crypto v0.39.0 // indirect
	golang.org/x/exp v0.0.0-20250620022241-b7579e27df2b // indirect
	golang.org/x/mod v0.25.0 // indirect
	golang.org/x/sync v0.15.0 // indirect
	golang.org/x/sys v0.33.0 // indirect
	golang.org/x/text v0.26.0 // indirect
	golang.org/x/time v0.7.0 // indirect
	golang.org/x/tools v0.34.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/loa/graphqlclientgen => ./../../

replace github.com/loa/graphqlclientgen/natsproto => ./../../natsproto/
