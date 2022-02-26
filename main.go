package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/serjbibox/FSTR/apicontroller"
	_ "github.com/serjbibox/FSTR/apicontroller"
	_ "github.com/serjbibox/FSTR/dbcontroller"
	_ "github.com/serjbibox/FSTR/docs"
	_ "github.com/serjbibox/FSTR/jsoncontroller"

	//_ "github.com/gorilla/mux"
	_ "github.com/lib/pq"
	//swaggerFiles "github.com/swaggo/files"
	//"github.com/swaggo/http-swagger"
)

const port = ":8080"

// @title Blueprint Swagger API
// @version 1.0
// @description Swagger API for Golang Project Blueprint.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email martin7.heinz@gmail.com

// @license.name MIT
// @license.url https://github.com/MartinHeinz/go-project-blueprint/blob/master/LICENSE

// @BasePath /api/v1
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
