package api

import (
	"fmt"
	"net/http"
)

func UpdateRecordHandler(w http.ResponseWriter, r *http.Request) {
	recordId := r.URL.Query().Get(recordIdKey)
	fmt.Fprintf(w, "RecordID obtained=%s\n", recordId)
}
