package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/serjbibox/FSTR/apis"
	"github.com/serjbibox/FSTR/dbcontroller"
	"github.com/serjbibox/FSTR/models"

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
		//r.With(paginate).Get("/", ListPass)
		r.Post("/", apis.Create)
		//r.Get("/search", SearchPass)

		r.Route("/{passID}", func(r chi.Router) {
			//r.Route("/:id", func(r chi.Router) {
			r.Use(PassCtx)
			r.Get("/", apis.GetPass)
			//r.Put("/", UpdateArticle)
			//r.Delete("/", DeleteArticle)
		})

		// GET /articles/whats-up
		//r.With(ArticleCtx).Get("/{articleSlug:[a-z-]+}", GetArticle)
	})

	//r.Post("/submitData", api.Create)
	//r.Get("/submitData/{id}", api.GetById)
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

type Pass string

func PassCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var p *models.Pereval
		var err error
		//if id := r.URL.Query().Get("passID"); id != "" {
		if id := chi.URLParam(r, "passID"); id != "" {
			if p, err = dbcontroller.GetRow(id); err != nil {
				apis.SendErr(w, r, http.StatusServiceUnavailable, fmt.Errorf("%w", err))
				return
			}
		} else {
			err = errors.New("URI not found")
			apis.SendErr(w, r, http.StatusServiceUnavailable, fmt.Errorf("%w", err))
			return
		}
		if err != nil {
			apis.SendErr(w, r, http.StatusServiceUnavailable, fmt.Errorf("%w", err))
			return
		}
		var ps Pass = "pass"
		ctx := context.WithValue(r.Context(), ps, p)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
