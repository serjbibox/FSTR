package api

import (
	"encoding/json"
	"net/http"

	"github.com/serjbibox/FSTR/dbcontroller"
	"github.com/serjbibox/FSTR/jsoncontroller"
)

// submitData godoc
// @Summary Retrieves user based on given ID
// @Produce json
// @Router /submitData [post]
func SubmitData(w http.ResponseWriter, r *http.Request) {
	var err error
	p := jsoncontroller.NewPereval()
	if err = json.NewDecoder(r.Body).Decode(&p); err != nil {
		SendErr(w, http.StatusServiceUnavailable, err)
		return
	}
	var id string
	if id, err = dbcontroller.AddData(&p); err != nil {
		SendErr(w, http.StatusServiceUnavailable, err)
		return
	}
	SendResponse(w, id)
}
