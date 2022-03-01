package main

import (
	"log"
	"net/http"

	"github.com/serjbibox/FSTR/apis"

	//"github.com/go-chi/chi"

	_ "github.com/serjbibox/FSTR/dbcontroller"

	_ "github.com/serjbibox/FSTR/models"

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
	//api := apis.Api{}
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))
	r.Get("/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8080/swagger/doc.json"))) // API definition
	r.Get("/panic", func(w http.ResponseWriter, r *http.Request) {
		panic("test")
	})

	r.Route("/submitData", func(r chi.Router) {
		//r.With(paginate).Get("/", ListPass)
		r.Post("/", apis.Create) // POST /articles
		//r.Get("/search", SearchPass) // GET /articles/search

		//r.Route("/{passID}", func(r chi.Router) {
		r.Route("/:id", func(r chi.Router) {
			r.Use(apis.PassCtx)      // Load the *Article on the request context
			r.Get("/", apis.GetPass) // GET /articles/123
			//r.Put("/", UpdateArticle)    // PUT /articles/123
			//r.Delete("/", DeleteArticle) // DELETE /articles/123
		})

		// GET /articles/whats-up
		//r.With(ArticleCtx).Get("/{articleSlug:[a-z-]+}", GetArticle)
	})
	//r.Get("/*", httpSwagger.Handler(httpSwagger.URL("http://propane-facet-342315.ue.r.appspot.com/swagger/doc.json"))) // API definition

	//r.Post("/submitData", api.Create)
	//r.Get("/submitData/{id}", api.GetById)
	log.Panic(http.ListenAndServe(port, r))
}

// paginate is a stub, but very possible to implement middleware logic
// to handle the request params for handling a paginated request.
func paginate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// just a stub.. some ideas are to look at URL query params for something like
		// the page number, or the limit, and send a query cursor down the chain
		next.ServeHTTP(w, r)
	})
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
