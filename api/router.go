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
	{Name: "Create record", Path: "/createRecord", Action: "POST", Handler: CreateRecordHandler},
	{Name: "Read record", Path: "/readRecord", Action: "GET", Handler: ReadRecordHandler},
	{Name: "Update record", Path: "/updateRecord", Action: "PUT", Handler: UpdateRecordHandler},
	{Name: "Delete record", Path: "/deleteRecord", Action: "DELETE", Handler: DeleteRecordHandler},
	{Name: "List records", Path: "/listRecords", Action: "POST", Handler: ListRecordsHandler},
}

func NewRouter() *mux.Router {
	newRouter := mux.NewRouter()
	for _, route := range routes {
		newRouter.HandleFunc(route.Path, route.Handler).Methods(route.Action)
	}
	return newRouter
}
