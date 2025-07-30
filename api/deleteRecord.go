package api

import (
	"common"
	"db"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func DeleteRecordHandler(w http.ResponseWriter, r *http.Request) {
	bbody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("couldnt read request body with err ", err.Error())
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

	err = db.DeleteRecord(newRequest.UserId, newRequest.ServiceName)
	if err != nil {
		fmt.Println("couldnt  delete record")
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte("record deleted"))
}
