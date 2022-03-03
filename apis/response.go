package apis

import (
	"encoding/json"
	"net/http"

	"github.com/serjbibox/FSTR/models"
)

type ResponseInterface interface {
	Send(w http.ResponseWriter)
}

// @Description Структура HTTP ответа:
// @Description если отправка успешна, дополнительно возвращается id вставленной записи.
type PassResponse struct {
	Pass *models.Pass
}

func (p PassResponse) Send(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&p.Pass)
}

// @Description Структура HTTP ответа:
// @Description если отправка успешна, дополнительно возвращается id вставленной записи.
type InsertResponse struct {
	ID      string `json:"id" example:"123"`
	Message string `json:"message" example:"OK"`
}

func (p InsertResponse) Send(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&p)
}

// @Description Структура HTTP ответа:
// @Description если отправка успешна, дополнительно возвращается id вставленной записи.
type StatusResponse struct {
	ID     string `json:"id" example:"123"`
	Status string `json:"status" example:"new"`
}

func (s StatusResponse) Send(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&s)
}

// @Description Структура HTTP ответа об ошибке
type ErrResponse struct {
	HTTPStatusCode int    `json:"-"`               // http response status code
	ErrorText      string `json:"error,omitempty"` // application-level error message, for debugging
}

func (err *ErrResponse) Send(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(err.HTTPStatusCode)
	json.NewEncoder(w).Encode(&err)
}
