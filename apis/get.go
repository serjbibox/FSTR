package apis

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

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

// @Summary  Получает запись из pereval_added по ID записи
// @Tags         /submitData/{id}
// @Produce      json
// @Param    id   path      int  true  "pereval_added PRIMARY KEY ID"
// @Success  200  {object}  apis.PassResponse
// @Failure  400  {object}  apis.ErrResponse
// @Failure  503  {object}  apis.ErrResponse
// @Router   /submitData/{id} [get]
func GetPass(w http.ResponseWriter, r *http.Request) {
	if id, ok := r.Context().Value("id").(string); !ok {
		err := errors.New("ошибка контекста GetPass")
		SendErr(w, http.StatusServiceUnavailable, err)
		return
	} else {
		if _, err := strconv.Atoi(id); err != nil {
			SendErr(w, http.StatusServiceUnavailable, err)
			return
		}
		s := services.New(daos.NewPassDAO())
		f := services.NewFlow()
		f.ID = id
		f.GetWith = singleMode
		if f := s.Get(f); f.Err != nil {
			SendErr(w, http.StatusServiceUnavailable, f.Err)
			return
		} else {
			SendHttp(w, PassResponse{PassLoaded: f.PassLoaded})
		}
	}

}

// @Summary  Получает записи из pereval_added по ID записи
// @Tags     /submitData/{id}/status
// @Produce  json
// @Param    id   path      int  true  "pereval_added PRIMARY KEY ID"
// @Success  200  {object}  apis.StatusResponse
// @Failure  400  {object}  apis.ErrResponse
// @Failure  503  {object}  apis.ErrResponse
// @Router   /submitData/{id}/status [get]
func GetStatus(w http.ResponseWriter, r *http.Request) {
	if id, ok := r.Context().Value("id").(string); !ok {
		err := errors.New("ошибка контекста GetPass")
		SendErr(w, http.StatusServiceUnavailable, err)
		return
	} else {
		if _, err := strconv.Atoi(id); err != nil {
			SendErr(w, http.StatusServiceUnavailable, err)
			return
		}
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

// @Summary      список всех данных для отображения, которые этот пользователь отправил на сервер через приложение с возможностью фильтрации по данным пользователя (ФИО, телефон, почта), если передан объект.
// @Description  К сожалению, Swagger считает, что у GET запроса не должно быть тела. Можно проверить этот запрос, например, в Postman.
// @Tags     /submitData/
// @Accept       json
// @Produce  json
// @Param        email  body      apis.QueryParam  false  "Фильтровать по email"  Format(email)
// @Success      200    {object}  apis.PassArrayResponse
// @Failure      400    {object}  apis.ErrResponse
// @Failure      503    {object}  apis.ErrResponse
// @Router       /submitData/ [get]
func ListPass(w http.ResponseWriter, r *http.Request) {

	q := QueryParam{}
	s := services.New(daos.NewPassDAO())
	f := services.NewFlow()
	if err := json.NewDecoder(r.Body).Decode(&q); err != nil {
		SendErr(w, http.StatusServiceUnavailable, err)
	}
	log.Println(q.Email, q.Phone)
	switch {
	case q.Email != "":
		f.GetBy = q.Email
		f.GetWith = mailMode
	case q.Phone != "":
		f.GetBy = q.Phone
		f.GetWith = phoneMode
	case q.Fam != "":
		f.GetByFIO[0] = q.Fam
		if q.Name == "" {
			err := errors.New("ошибка, отсутствует параметр \"name\"")
			SendErr(w, http.StatusServiceUnavailable, err)
			return
		}
		f.GetByFIO[1] = q.Name
		if q.Otc == "" {
			err := errors.New("ошибка, отсутствует параметр \"otc\"")
			SendErr(w, http.StatusServiceUnavailable, err)
			return
		}
		f.GetByFIO[2] = q.Otc
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
