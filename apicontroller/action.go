package apicontroller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/serjbibox/FSTR/dbcontroller"
	"github.com/serjbibox/FSTR/jsoncontroller"
)

func SubmitData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p := jsoncontroller.NewPereval()
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		resp := make(map[string]string)
		resp["message"] = fmt.Sprintf("ошибка: %s", err)
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			log.Println(err)
		}
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write(jsonResp)
		log.Println(err)
		return
	}
	if id, err := dbcontroller.AddData(&p); err != nil {
		resp := make(map[string]string)
		resp["message"] = fmt.Sprintf("ошибка: %s", err)
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			log.Println(err)
		}
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write(jsonResp)
		log.Println(err)
		return
	} else {
		resp := make(map[string]string)
		resp["message"] = "OK, ID: " + id
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			log.Println(err)
		}
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResp)
	}

}
