package apis

import (
	"errors"
	"net/http"

	"github.com/serjbibox/FSTR/daos"
	"github.com/serjbibox/FSTR/services"
)

func UpdatePass(w http.ResponseWriter, r *http.Request) {
	var replaceId string
	if id, ok := r.Context().Value("id").(string); !ok {
		err := errors.New("ошибка контекста GetPass")
		SendErr(w, http.StatusServiceUnavailable, err)
		return
	} else {
		replaceId = id
	}
	s := services.New(daos.NewPassDAO())
	f := services.NewFlow()
	f.ErrStatus = http.StatusServiceUnavailable
	f.ID = replaceId
	f.RID = ""
	f = s.Get(f)
	if f.Err != nil {
		SendErr(w, f.ErrStatus, f.Err)
		return
	}
	if f.Pass.Status != new {
		f.Err = errors.New("статус объекта: " + f.Pass.Status)
		SendErr(w, f.ErrStatus, f.Err)
		return
	}
	storedPass := f.Pass
	f = services.NewFlow()
	f.Pass = s.Create(f, r).
		ValidateFields().
		ValidateData().
		GetImage().
		Pass
	if f.Err != nil {
		SendErr(w, f.ErrStatus, f.Err)
		return
	}
	f.Pass.User.Email = storedPass.User.Email
	f.Pass.User.Name = storedPass.User.Name
	f.Pass.User.Fam = storedPass.User.Fam
	f.Pass.User.Otc = storedPass.User.Otc
	f.Pass.User.Phone = storedPass.User.Phone
	f.RID = replaceId
	s.InsertTo(f, "pereval_images").ImgData()
	s.InsertTo(f, "pereval_added")
	if f.Err != nil {
		SendErr(w, f.ErrStatus, f.Err)
		return
	}
	SendHttp(w,
		InsertResponse{
			Message: "OK",
			ID:      f.ID,
		})

}
