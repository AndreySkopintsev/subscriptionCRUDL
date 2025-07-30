package api

import (
	"common"
	"db"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	recordIdKey = "recordId"
)

// ReadRecordHandler - gets a record from db by UUID
func ReadRecordHandler(w http.ResponseWriter, r *http.Request) {
	bbody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("couldnt read request body")
		w.Write([]byte(err.Error()))
		return
	}
	newRequest := common.Subscription{}
	err = json.Unmarshal(bbody, &newRequest)
	if err != nil {
		fmt.Println("couldnt unmarshal request body")
		w.Write([]byte(err.Error()))
		return
	}
	subscription, err := db.GetRecord(newRequest.UserId, newRequest.ServiceName)
	if err != nil {
		fmt.Println("couldnt insert into db")
		w.Write([]byte(err.Error()))
		return
	}
	bbody, err = json.Marshal(subscription)
	if err != nil {
		fmt.Println("couldnt  marshal result")
		w.Write([]byte(err.Error()))
		return
	}
	w.Write(bbody)
}
