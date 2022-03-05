package apis

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/serjbibox/FSTR/daos"
	"github.com/serjbibox/FSTR/services"
)

// @Summary  Редактирует существующую запись в pereval_added , если она в статусе "new"
// @Tags     /submitData/{id}
// @Accept   json
// @Produce  json
// @Param    id     path      int          true  "pereval_added PRIMARY KEY ID"
// @Param    input  body      models.Pass  true  "карточка объекта"
// @Success  200    {object}  apis.InsertResponse
// @Failure  400    {object}  apis.ErrResponse
// @Failure  503    {object}  apis.ErrResponse
// @Router   /submitData/{id} [put]
func UpdatePass(w http.ResponseWriter, r *http.Request) {
	var replaceId string
	if id, ok := r.Context().Value("id").(string); !ok {
		err := errors.New("ошибка контекста GetPass")
		SendErr(w, http.StatusServiceUnavailable, err)
		return
	} else {
		if _, err := strconv.Atoi(id); err != nil {
			SendErr(w, http.StatusServiceUnavailable, err)
			return
		}
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
