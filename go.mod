module main

go 1.24.4

require api v0.0.0-00010101000000-000000000000

require db v0.0.0-00010101000000-000000000000

require common v0.0.0-00010101000000-000000000000 // indirect

require (
	github.com/gorilla/mux v1.8.1 // indirect
	github.com/lib/pq v1.10.9 // indirect
)

replace api => ./api

replace db => ./db

replace common => ./common
