package apis

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/serjbibox/FSTR/models"

	"github.com/go-chi/render"
)

func GetPass(w http.ResponseWriter, r *http.Request) {
	if p, ok := r.Context().Value("pass").(*models.Pereval); !ok {
		err := errors.New("ошибка контекста")
		SendErr(w, r, http.StatusServiceUnavailable, fmt.Errorf("%w", err))
		return
	} else {
		if err := render.Render(w, r, &PerevalResponse{
			Pereval: p,
		}); err != nil {
			SendErr(w, r, http.StatusServiceUnavailable, fmt.Errorf("%w", err))
			return
		}
	}
}

type PerevalResponse struct {
	*models.Pereval
}

func (rd *PerevalResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
