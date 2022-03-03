package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/serjbibox/FSTR/apis"

	_ "github.com/lib/pq"
	_ "github.com/serjbibox/FSTR/docs"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	httpSwagger "github.com/swaggo/http-swagger"
)

const port = ":8080"

// @title          ФСТР API
// @version        1.0
// @description    API для взаимодействия приложения с сервером БД ФСТР.
// @contact.name   API Support
// @contact.email  serj_bibox@mail.ru

// @BasePath
func main() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))
	// API definition localhost
	r.Get("/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8080/swagger/doc.json")))
	// API definition
	//r.Get("/*", httpSwagger.Handler(httpSwagger.URL("http://propane-facet-342315.ue.r.appspot.com/swagger/doc.json")))
	r.Get("/panic", func(w http.ResponseWriter, r *http.Request) {
		panic("test")
	})

	r.Route("/submitData", func(r chi.Router) {
		r.Get("/", apis.Filter)
		r.Post("/", apis.Insert)
		r.Route("/{passID}", func(r chi.Router) {
			r.Use(Ctx)
			r.Get("/", apis.GetPass)
			r.Get("/status", apis.GetStatus)
			r.Put("/", apis.UpdatePass)
		})
	})
	log.Panic(http.ListenAndServe(port, r))
}

/*
GET /submitData/:id/status — получить статус модерации отправленных данных. OK
PUT /submitData/:id — отредактировать существующую запись (замена), если она в статусе new.
Редактировать можно все поля, кроме ФИО, почта, телефон. OK
GET /submitData/ — список всех данных для отображения, которые этот пользователь отправил
на сервер через приложение с возможностью фильтрации по данным пользователя (ФИО, телефон, почта), если передан объект.
GET /submitData/:id — получить одну запись (перевал) по её id. OK
При создании записи в БД, бэк возвращает фронту id и фронт этот id сохраняет у себя локально.
За счёт этого может редактировать записи, которые ещё не отрезолвлены модератором.
*/

func Ctx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if id := chi.URLParam(r, "passID"); id != "" {
			ctx := context.WithValue(r.Context(), "id", id)
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			err := errors.New("URI not found")
			apis.SendErr(w, http.StatusServiceUnavailable, fmt.Errorf("%w", err))
			return
		}
	})
}
