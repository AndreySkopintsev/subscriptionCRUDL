module main

go 1.24.4

require api v0.0.0-00010101000000-000000000000

require db v0.0.0-00010101000000-000000000000

require common v0.0.0 // indirect

require (
	github.com/golang-migrate/migrate/v4 v4.18.3 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/gorilla/mux v1.8.1 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/lib/pq v1.10.9 // indirect
	go.uber.org/atomic v1.7.0 // indirect
)

replace api => ./api

replace db => ./db

replace common => ./common
