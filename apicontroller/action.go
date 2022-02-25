package apicontroller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/serjbibox/FSTR/dbcontroller"
	"github.com/serjbibox/FSTR/jsoncontroller"
)

func SubmitData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p := jsoncontroller.NewPereval()
	resp := jsoncontroller.NewResponse()
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		resp.Status = 503
		resp.Message = "ошибка: " + err.Error()
		json.NewEncoder(w).Encode(&resp)
		log.Println(err)
	}
	if id, err := dbcontroller.AddData(&p); err != nil {
		resp.Status = 503
		resp.Message = "ошибка: " + err.Error()
		json.NewEncoder(w).Encode(&resp)
	} else {
		resp.Status = 200
		resp.Message = "ok, ID: " + id
		json.NewEncoder(w).Encode(&resp)
	}

}
