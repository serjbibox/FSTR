package apis

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/go-chi/render"
)

// @Summary   Получает запись из pereval_added по ID записи
// @Tags /submitData/:id
// @Produce   json
// @Success   200  {object}  apis.Response
// @Failure   400  {object}  apis.ErrResponse
// @Failure   503  {object}  apis.ErrResponse
// @Router    /submitData/:id [get]
func GetPass(w http.ResponseWriter, r *http.Request) {
	if ctx, ok := r.Context().Value("pass").(*Context); !ok {
		err := errors.New("ошибка контекста GetPass")
		SendErr(w, r, http.StatusServiceUnavailable, fmt.Errorf("%w", err))
		return
	} else {
		if err := render.Render(w, r, &PerevalResponse{
			Pereval: ctx.Pereval,
		}); err != nil {
			SendErr(w, r, http.StatusServiceUnavailable, fmt.Errorf("%w", err))
			return
		}
	}
}

func GetStatus(w http.ResponseWriter, r *http.Request) {
	if ctx, ok := r.Context().Value("pass").(*Context); !ok {
		err := errors.New("ошибка контекста GetStatus")
		SendErr(w, r, http.StatusServiceUnavailable, fmt.Errorf("%w", err))
		return
	} else {
		if err := render.Render(w, r, &PerevalResponse{Status: *ctx.Status}); err != nil {
			SendErr(w, r, http.StatusServiceUnavailable, fmt.Errorf("%w", err))
			return
		}
	}

}
