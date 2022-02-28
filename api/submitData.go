package api

import (
	"encoding/json"
	"net/http"

	"github.com/serjbibox/FSTR/dbcontroller"
	"github.com/serjbibox/FSTR/models"
)

// submitData godoc
// @Summary Retrieves user based on given ID
// @Produce json
// @Router /submitData [post]
func SubmitData(w http.ResponseWriter, r *http.Request) {

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
