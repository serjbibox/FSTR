package apis

import (
	"net/http"

	"github.com/serjbibox/FSTR/daos"
	"github.com/serjbibox/FSTR/services"
)

// @Summary  Создаёт новую запись в pereval_added
// @Tags     /submitData
// @Accept   json
// @Produce  json
// @Param    input  body      models.Pass  true  "карточка объекта"
// @Success  200    {object}  apis.InsertResponse
// @Failure  400    {object}  apis.ErrResponse
// @Failure  503    {object}  apis.ErrResponse
// @Router   /submitData [post]
func Insert(w http.ResponseWriter, r *http.Request) {
	s := services.New(daos.NewPassDAO())
	f := services.NewFlow()
	f.ErrStatus = http.StatusServiceUnavailable
	f.Pass = s.Create(f, r).
		ValidateFields().
		ValidateData().
		GetImage().
		Pass
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
