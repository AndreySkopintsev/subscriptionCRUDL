module api

go 1.24.4

require github.com/gorilla/mux v1.8.1

require common v0.0.0

require db v0.0.0-00010101000000-000000000000

require (
	github.com/golang-migrate/migrate/v4 v4.18.3 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/lib/pq v1.10.9 // indirect
	go.uber.org/atomic v1.7.0 // indirect
)

replace common => ../common

replace db => ../db
