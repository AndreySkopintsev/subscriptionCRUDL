package api

import (
	"fmt"
	"net/http"
)

const (
	recordIdKey = "recordId"
)

func ReadRecordHandler(w http.ResponseWriter, r *http.Request) {
	recordId := r.URL.Query().Get(recordIdKey)
	fmt.Fprintf(w, "RecordID obtained=%s\n", recordId)
}
