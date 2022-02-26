package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/serjbibox/FSTR/apicontroller"
	_ "github.com/serjbibox/FSTR/apicontroller"
	_ "github.com/serjbibox/FSTR/dbcontroller"
	_ "github.com/serjbibox/FSTR/jsoncontroller"

	//_ "github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const port = ":8080"

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/submitData", apicontroller.SubmitData).Methods("POST")
	log.Panic(http.ListenAndServe(port, r))
}

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
