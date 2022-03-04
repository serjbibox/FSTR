package apis

import (
	"net/http"

	"github.com/serjbibox/FSTR/daos"
	"github.com/serjbibox/FSTR/services"
)

// submitData godoc
// @Summary   Создаёт новую запись в pereval_added
// @Tags /submitData
// @Accept    json
// @Produce   json
// @Param     input	body	models.Pereval true "карточка объекта"
// @Param     output	body	apis.Response true "ID созданной записи"
// @Success   200  {object}  apis.Response
// @Failure   400  {object}  apis.ErrResponse
// @Failure   503  {object}  apis.ErrResponse
// @Router    /submitData [post]
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
