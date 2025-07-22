package api

import (
	"fmt"
	"net/http"
)

func DeleteRecordHandler(w http.ResponseWriter, r *http.Request) {
	recordId := r.URL.Query().Get(recordIdKey)
	fmt.Fprintf(w, "RecordID obtained=%s\n", recordId)
}
