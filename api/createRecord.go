package api

import (
	"common"
	"db"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func CreateRecordHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("couldnt read request body")
		w.Write([]byte(err.Error()))
		return
	}
	requestStruct := common.Subscription{}
	err = json.Unmarshal(body, &requestStruct)
	if err != nil {
		fmt.Println("couldnt unmarshal request body")
		w.Write([]byte(err.Error()))
		return
	}
	_, err = db.CreateNewRecord(requestStruct.Price, requestStruct.ServiceName)
	if err != nil {
		fmt.Println("couldnt insert into db")
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte("it worked, you have my word"))
}
