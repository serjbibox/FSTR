package apis

import (
	"errors"
	"net/http"

	"github.com/serjbibox/FSTR/daos"
	"github.com/serjbibox/FSTR/services"
)

// @Summary   Получает запись из pereval_added по ID записи
// @Tags /submitData/:id
// @Produce   json
// @Success   200  {object}  apis.Response
// @Failure   400  {object}  apis.ErrResponse
// @Failure   503  {object}  apis.ErrResponse
// @Router    /submitData/:id [get]
func GetPass(w http.ResponseWriter, r *http.Request) {
	if id, ok := r.Context().Value("id").(string); !ok {
		err := errors.New("ошибка контекста GetPass")
		SendErr(w, http.StatusServiceUnavailable, err)
		return
	} else {
		s := services.New(daos.NewPassDAO())
		if p, err := s.Get(id); err != nil {
			SendErr(w, http.StatusServiceUnavailable, err)
			return
		} else {
			SendHttp(w, PassResponse{Pass: p})
		}
	}

}

func GetStatus(w http.ResponseWriter, r *http.Request) {
	if id, ok := r.Context().Value("id").(string); !ok {
		err := errors.New("ошибка контекста GetPass")
		SendErr(w, http.StatusServiceUnavailable, err)
		return
	} else {
		s := services.New(daos.NewPassDAO())
		if status, err := s.GetStatus(id); err != nil {
			SendErr(w, http.StatusServiceUnavailable, err)
			return
		} else {
			SendHttp(w, StatusResponse{
				ID:     id,
				Status: status,
			})
		}
	}
}
