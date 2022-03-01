package apis

import (
	"encoding/json"
	"net/http"

	"github.com/serjbibox/FSTR/dbcontroller"
	"github.com/serjbibox/FSTR/models"
)

// submitData godoc
// @Summary   Создаёт новую запись в pereval_added
// @Tags /submitData
// @Accept    json
// @Produce   json
// @Param     input	body	models.Pereval true "карточка объекта"
// @Success   200  {object}  models.Response
// @Failure   400  {object}  models.ResponseErr
// @Failure   503  {object}  models.ResponseErr
// @Router    /submitData [post]
func Create(w http.ResponseWriter, r *http.Request) {
	p := models.NewPereval()
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		SendErr(w, http.StatusServiceUnavailable, err)
		return
	}
	if err := p.ValidateFields(); err != nil {
		SendErr(w, http.StatusBadRequest, err)
		return
	}
	if err := p.ValidateData(); err != nil {
		SendErr(w, http.StatusServiceUnavailable, err)
		return
	}
	if id, err := dbcontroller.AddData(&p); err != nil {
		SendErr(w, http.StatusServiceUnavailable, err)
		return
	} else {
		SendResponse(w, id)
	}

}
