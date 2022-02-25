package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	_ "github.com/serjbibox/FSTR/apicontroller"
	"github.com/serjbibox/FSTR/dbcontroller"
	"github.com/serjbibox/FSTR/jsoncontroller"

	"github.com/gorilla/mux"
	//_ "github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

/*
GET /submitData/:id/status — получить статус модерации отправленных данных.
PUT /submitData/:id — отредактировать существующую запись (замена), если она в статусе new.
Редактировать можно все поля, кроме ФИО, почта, телефон.
GET /submitData/ — список всех данных для отображения, которые этот пользователь отправил
на сервер через приложение с возможностью фильтрации по данным пользователя (ФИО, телефон, почта), если передан объект.
GET /submitData/:id — получить одну запись (перевал) по её id.
При создании записи в БД, бэк возвращает фронту id и фронт этот id сохраняет у себя локально.
За счёт этого может редактировать записи, которые ещё не отрезолвлены модератором.
*/
const (
	new = "new"

	pending  = "pending"
	resolved = "resolved"
	accepted = "accepted"
	rejected = "rejected"
)

const port = ":85"

var Status string = new

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/submitData", submitData).Methods("POST")
	log.Panic(http.ListenAndServe(port, r))
}

func submitData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p := jsoncontroller.NewPereval()
	resp := jsoncontroller.NewResponse()
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		resp.Status = 503
		resp.Message = "ошибка: " + err.Error()
		json.NewEncoder(w).Encode(&resp)
		log.Println(err)
	}
	if id, err := addData(&p); err != nil {
		resp.Status = 503
		resp.Message = "ошибка: " + err.Error()
		json.NewEncoder(w).Encode(&resp)
	} else {
		resp.Status = 200
		resp.Message = "ok, ID: " + id
		json.NewEncoder(w).Encode(&resp)
	}

}

func addData(p *jsoncontroller.Pereval) (id string, err error) {
	var t time.Time
	dbcontroller.Connect()
	defer dbcontroller.DB.Close()
	if t, err = time.Parse("2006-01-02 15:04:05", p.AddTime); err != nil {
		log.Println(err)
		return "", err
	}
	data, err := json.Marshal(p)
	if err != nil {
		log.Println(err)
		return "", err
	}
	Status = new
	err = dbcontroller.DB.QueryRow("INSERT INTO pereval_added (date_added, raw_data, status) VALUES ($1, $2, $3) RETURNING ID;", t, data, Status).Scan(&id)
	//result, err := db.Exec("INSERT INTO pereval_added (date_added, raw_data, status) VALUES ($1, $2, $3);", t, data, Status)
	if err != nil {
		log.Println(err)
		return "", err
	}
	log.Println(id)
	return id, nil
}
