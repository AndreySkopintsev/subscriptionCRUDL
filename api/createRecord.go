package api

import "net/http"

func CreateRecordHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("there will be some fucntionality, i swear"))
}
