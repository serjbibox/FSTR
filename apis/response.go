package apis

import (
	"encoding/json"
	"net/http"

	"github.com/serjbibox/FSTR/models"
)

type ResponseInterface interface {
	Send(w http.ResponseWriter)
}

// @Description Структура HTTP ответа метода GET /submitData/{id}
type PassResponse struct {
	PassLoaded *models.PassLoaded
}

func (p PassResponse) Send(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&p.PassLoaded)
}

// @Description Структура HTTP ответа метода PUT /submitData/{id}
type PassArrayResponse struct {
	Parray *[]models.PassLoaded
}

func (p PassArrayResponse) Send(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&p.Parray)
}

// @Description Структура HTTP ответа метода POST /submitData
type InsertResponse struct {
	ID      string `json:"id" example:"123"`
	Message string `json:"message" example:"OK"`
}

func (p InsertResponse) Send(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&p)
}

// @Description Структура HTTP ответа метода GET /submitData/{id}/status
type StatusResponse struct {
	ID     string `json:"id" example:"123"`
	Status string `json:"status" example:"new"`
}

func (s StatusResponse) Send(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&s)
}

// @Description Структура HTTP ответа ERROR
type ErrResponse struct {
	HTTPStatusCode int    `json:"-"`
	ErrorText      string `json:"error,omitempty"`
}

func (err *ErrResponse) Send(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(err.HTTPStatusCode)
	json.NewEncoder(w).Encode(&err)
}

// @Description Структура тела запроса метода GET /submitData Необходимые поля: Phone | Email | Fam & Name & OTC
type QueryParam struct {
	Phone string `json:"phone" example:"+71234567890"`
	Email string `json:"email" example:"sample@sample.com"`
	Fam   string `json:"fam" example:"Иванов"`
	Name  string `json:"name" example:"Иван"`
	Otc   string `json:"otc" example:"Иванович"`
}
