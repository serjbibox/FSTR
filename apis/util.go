package apis

import (
	"fmt"
	"net/http"

	"github.com/go-chi/render"
)

func SendErr(w http.ResponseWriter, r *http.Request, statusCode int, err error) {
	render.Render(w, r, &ErrResponse{
		HTTPStatusCode: statusCode,
		StatusText:     fmt.Sprintf("ошибка: %s", err),
	})
}

//var ErrNotFound = &ErrResponse{HTTPStatusCode: 404, StatusText: "Resource not found."}
