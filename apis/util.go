package apis

import (
	"fmt"
	"net/http"

	"github.com/go-chi/render"
	"github.com/serjbibox/FSTR/models"
)

func SendErr(w http.ResponseWriter, r *http.Request, statusCode int, err error) {
	render.Render(w, r, &ErrResponse{
		HTTPStatusCode: statusCode,
		StatusText:     fmt.Sprintf("ошибка: %s", err),
	})
}

type Context struct {
	Pereval *models.Pereval
	Status  *StatusResponse
}
