package apis

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/serjbibox/FSTR/dbcontroller"

	"github.com/go-chi/render"
	"github.com/serjbibox/FSTR/models"
)

func SendResponse(w http.ResponseWriter, data string) {
	//resp := make(map[string]string)
	//resp["message"] = "OK, ID: " + data
	resp := models.Response{
		Message: "OK",
		ID:      data,
	}
	jsonResp, _ := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResp)
}

func SendErr(w http.ResponseWriter, statusCode int, err error) {
	resp := make(map[string]string)
	resp["message"] = fmt.Sprintf("ошибка: %s", err)
	jsonResp, _ := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(jsonResp)
}

var ErrNotFound = &ErrResponse{HTTPStatusCode: 404, StatusText: "Resource not found."}

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
func ErrRender(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 422,
		StatusText:     "Error rendering response.",
		ErrorText:      err.Error(),
	}
}

func DbGetPass(id string) (*Pass, error) {
	for _, a := range PassRows {
		if a.ID == id {
			return a, nil
		}
	}
	return nil, errors.New("article not found")
}

func PassCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var p *models.Pereval
		var err error
		if id := r.URL.Query().Get("passID"); id != "" {
			p, err = dbcontroller.GetRow(id)
			//log.Println(p, err)
		} else {
			render.Render(w, r, ErrNotFound)
			return
		}
		if err != nil {
			render.Render(w, r, ErrNotFound)
			return
		}
		ctx := context.WithValue(r.Context(), "pass", p)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
