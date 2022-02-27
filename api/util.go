package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func SendResponse(w http.ResponseWriter, data string) {
	resp := make(map[string]string)
	resp["message"] = "OK, ID: " + data
	jsonResp, _ := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResp)
}

func SendErr(w http.ResponseWriter, statusCode int, err error) {
	resp := make(map[string]string)
	resp["message"] = fmt.Sprintf("ошибка: %s", err)
	jsonResp, _ := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(jsonResp)
}
