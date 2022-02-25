package apicontroller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/serjbibox/FSTR/dbcontroller"
	"github.com/serjbibox/FSTR/jsoncontroller"
)

func SubmitData(w http.ResponseWriter, r *http.Request) {
	var err error
	p := jsoncontroller.NewPereval()
	if err = json.NewDecoder(r.Body).Decode(&p); err != nil {
		sendResponseErr(w, err)
		return
	}
	var id string
	if id, err = dbcontroller.AddData(&p); err != nil {
		sendResponseErr(w, err)
		return
	}
	sendResponse(w, id)
}

func sendResponse(w http.ResponseWriter, data string) {
	resp := make(map[string]string)
	resp["message"] = "OK, ID: " + data
	jsonResp, _ := json.Marshal(resp)
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResp)
}

func sendResponseErr(w http.ResponseWriter, err error) {
	resp := make(map[string]string)
	w.Header().Set("Content-Type", "application/json")
	resp["message"] = fmt.Sprintf("ошибка: %s", err)
	jsonResp, _ := json.Marshal(resp)
	w.WriteHeader(http.StatusServiceUnavailable)
	w.Write(jsonResp)
}
