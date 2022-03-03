package apis

import (
	"net/http"

	"github.com/serjbibox/FSTR/daos"
	"github.com/serjbibox/FSTR/models"
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
	var err error
	var p *models.Pass
	s := services.New(daos.NewPassDAO())
	p, err = s.Create(r)
	if err != nil {
		SendErr(w, http.StatusServiceUnavailable, err)
		return
	}
	if err = Validate(p, s); err != nil {
		SendErr(w, http.StatusBadRequest, err)
		return
	}

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

	if id, err := s.Insert(p, imgMap, ""); err != nil {
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
