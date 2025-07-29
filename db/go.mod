module db

go 1.24.4

require github.com/lib/pq v1.10.9

require (
	common v0.0.0-00010101000000-000000000000
	github.com/golang-migrate/migrate/v4 v4.18.3
	github.com/google/uuid v1.6.0
)

require (
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	go.uber.org/atomic v1.7.0 // indirect
)

replace common => ../common
