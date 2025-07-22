package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name    string
	Path    string
	Action  string
	Handler func(w http.ResponseWriter, r *http.Request)
}

// TODO add handlers
var routes []Route = []Route{
	{Name: "Create record", Path: "/createRecord", Action: "POST", Handler: func(w http.ResponseWriter, r *http.Request) {}},
	{Name: "Read record", Path: "/readRecord", Action: "GET", Handler: func(w http.ResponseWriter, r *http.Request) {}},
	{Name: "Update record", Path: "/updateRecord", Action: "PUT", Handler: func(w http.ResponseWriter, r *http.Request) {}},
	{Name: "Delete record", Path: "/deleteRecord", Action: "DELETE", Handler: func(w http.ResponseWriter, r *http.Request) {}},
	{Name: "List records", Path: "/listRecords", Action: "POST", Handler: func(w http.ResponseWriter, r *http.Request) {}},
}

func NewRouter() *mux.Router {
	newRouter := mux.NewRouter()
	for _, route := range routes {
		newRouter.HandleFunc(route.Path, route.Handler).Methods(route.Action)
	}
	return newRouter
}
