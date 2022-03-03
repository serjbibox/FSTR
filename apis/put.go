package apis

import (
	"errors"
	"net/http"

	"github.com/serjbibox/FSTR/daos"
	"github.com/serjbibox/FSTR/models"
	"github.com/serjbibox/FSTR/services"
)

func UpdatePass(w http.ResponseWriter, r *http.Request) {
	var err error
	var replaceId string
	if id, ok := r.Context().Value("id").(string); !ok {
		err := errors.New("ошибка контекста GetPass")
		SendErr(w, http.StatusServiceUnavailable, err)
		return
	} else {
		replaceId = id
	}
	s := services.New(daos.NewPassDAO())
	var prev *models.Pass
	if prev, err = s.Get(replaceId); err != nil {
		SendErr(w, http.StatusServiceUnavailable, err)
		return
	}

	var p *models.Pass
	p, err = s.Create(r)
	if err != nil {
		SendErr(w, http.StatusServiceUnavailable, err)
		return
	}
	if err = Validate(p, s); err != nil {
		SendErr(w, http.StatusBadRequest, err)
		return
	}

	p.User.Email = prev.User.Email
	p.User.Name = prev.User.Name
	p.User.Fam = prev.User.Fam
	p.User.Otc = prev.User.Otc
	p.User.Phone = prev.User.Phone

	var img [][]byte
	if img, err = GetImage(p); err != nil {
		SendErr(w, http.StatusServiceUnavailable, err)
		return
	}
	var m map[string]string
	if m, err = s.InsertImage(p, img); err != nil {
		SendErr(w, http.StatusServiceUnavailable, err)
		return
	}
	var imgMap *map[string][]int
	if imgMap, err = imgData(m); err != nil {
		SendErr(w, http.StatusServiceUnavailable, err)
		return
	}
	if id, err := s.Insert(p, imgMap, replaceId); err != nil {
		SendErr(w, http.StatusServiceUnavailable, err)
		return
	} else {
		SendHttp(w,
			InsertResponse{
				Message: "OK",
				ID:      id,
			})
	}

}
