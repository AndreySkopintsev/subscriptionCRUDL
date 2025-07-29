module api

go 1.24.4

require github.com/gorilla/mux v1.8.1
require common v0.0.0
require db v0.0.0-00010101000000-000000000000

replace common => ../common
replace db => ../db