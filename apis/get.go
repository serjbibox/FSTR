package apis

import (
	"errors"
	"net/http"

	"github.com/serjbibox/FSTR/daos"
	"github.com/serjbibox/FSTR/services"
)

const (
	singleMode = iota
	mailMode
	phoneMode
	fioMode
	statusMode
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
		f := services.NewFlow()
		f.ID = id
		f.GetWith = singleMode
		if f := s.Get(f); f.Err != nil {
			SendErr(w, http.StatusServiceUnavailable, f.Err)
			return
		} else {
			SendHttp(w, PassResponse{Pass: f.Pass})
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
		f := services.NewFlow()
		f.ID = id
		f.GetWith = statusMode
		if f = s.Get(f); f.Err != nil {
			SendErr(w, http.StatusServiceUnavailable, f.Err)
			return
		}
		SendHttp(w, StatusResponse{
			ID:     id,
			Status: f.Pass.Status,
		})
	}
}

func ListPass(w http.ResponseWriter, r *http.Request) {
	s := services.New(daos.NewPassDAO())
	f := services.NewFlow()
	switch {
	case r.URL.Query().Get("email") != "":
		f.GetBy = r.URL.Query().Get("email")
		f.GetWith = mailMode
	case r.URL.Query().Get("phone") != "":
		f.GetBy = r.URL.Query().Get("phone")
		f.GetWith = phoneMode
	case r.URL.Query().Get("fam") != "":
		f.GetByFIO[0] = r.URL.Query().Get("fam")
		if r.URL.Query().Get("name") == "" {
			err := errors.New("ошибка, отсутствует параметр \"name\"")
			SendErr(w, http.StatusServiceUnavailable, err)
			return
		}
		f.GetByFIO[1] = r.URL.Query().Get("name")
		if r.URL.Query().Get("otc") == "" {
			err := errors.New("ошибка, отсутствует параметр \"otc\"")
			SendErr(w, http.StatusServiceUnavailable, err)
			return
		}
		f.GetByFIO[2] = r.URL.Query().Get("otc")
		f.GetWith = fioMode
	default:
		err := errors.New("ошибка, отсутствует параметр")
		SendErr(w, http.StatusServiceUnavailable, err)
		return
	}
	if f := s.Get(f); f.Err != nil {
		SendErr(w, http.StatusServiceUnavailable, f.Err)
		return
	} else {
		SendHttp(w, PassArrayResponse{Parray: &f.Parray})
	}

}
