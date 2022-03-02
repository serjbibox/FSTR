package apis

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/render"
	"github.com/serjbibox/FSTR/models"
)

// @Description Структура HTTP ответа:
// @Description если отправка успешна, дополнительно возвращается id вставленной записи.
type Response struct {
	Message string `json:"message" example:"OK"`
	ID      string `json:"id" example:"123"`
}

// @Description Структура HTTP ответа:
// @Description если отправка успешна, дополнительно возвращается id вставленной записи.
type StatusResponse struct {
	ID     string `json:"id" example:"123"`
	Status string `json:"status" example:"new"`
}

// @Description Структура HTTP ответа об ошибке
type ErrResponse struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	StatusText string `json:"status"`          // user-level status message
	AppCode    int64  `json:"code,omitempty"`  // application-specific error code
	ErrorText  string `json:"error,omitempty"` // application-level error message, for debugging
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func SendResponse(w http.ResponseWriter, data string) {
	resp := Response{
		Message: "OK",
		ID:      data,
	}
	jsonResp, _ := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResp)
}

type PerevalResponse struct {
	*models.Pereval
	*StatusResponse
	*Response
	Status StatusResponse
}

func (rd *PerevalResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
