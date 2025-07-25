package main

import (
	"api"
	"db"
	"fmt"
	"net/http"
)

func main() {
	newRouter := api.NewRouter()
	db.Init()
	err := http.ListenAndServe(":8080", newRouter)
	if err != nil {
		fmt.Printf("encountered error while listeningon port 8080: %s", err.Error())
	}
}
