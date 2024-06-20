module github.com/353solutions/go-cookbook

go 1.21.0

toolchain go1.21.4

// Hack to make files example work
replace git.corp.com/client => ./funcs/files/client

require (
	git.corp.com/client v0.0.0-00010101000000-000000000000
	github.com/ardanlabs/conf/v3 v3.1.6
	github.com/google/uuid v1.3.1
	github.com/gorilla/mux v1.8.0
	github.com/mitchellh/mapstructure v1.5.0
	github.com/rs/xid v1.5.0
	github.com/stretchr/testify v1.8.4
	go.uber.org/zap v1.25.0
	golang.org/x/image v0.14.0
	golang.org/x/net v0.14.0
	golang.org/x/sys v0.11.0
	golang.org/x/text v0.14.0
	golang.org/x/tools v0.12.0
	gopkg.in/yaml.v2 v2.4.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	go.uber.org/multierr v1.10.0 // indirect
	golang.org/x/mod v0.12.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
