package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/serjbibox/FSTR/api"
	_ "github.com/serjbibox/FSTR/dbcontroller"

	_ "github.com/serjbibox/FSTR/models"

	_ "github.com/lib/pq"
	_ "github.com/serjbibox/FSTR/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

const port = ":8080"

// @title ФСТР API
// @version 1.0
// @description API для взаимодействия приложения с сервером БД ФСТР.
// @contact.name API Support
// @contact.email serj_bibox@mail.ru

// @BasePath /api/v1
func main() {
	/*var s string
	p := models.NewPereval()
	p.AddTime = "2021-09-22 13:18:13"
	if t, err := time.Parse("2006-01-02 15:04:05", p.AddTime); err != nil {
		log.Println(err)
		//return fmt.Errorf("%w", err)
	} else {
		log.Println(t.Format("2006-01-02 15:04:05"))
		s = t.String()
		log.Println(s)
	}*/

	r := chi.NewRouter()
	//r.Get("/*", httpSwagger.Handler(httpSwagger.URL("http://propane-facet-342315.ue.r.appspot.com/swagger/doc.json"))) // API definition
	r.Get("/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8080/swagger/doc.json"))) // API definition
	r.Post("/submitData", api.SubmitData)
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
