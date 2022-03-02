package apis

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/serjbibox/FSTR/dbcontroller"
	"github.com/serjbibox/FSTR/models"
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
func Create(w http.ResponseWriter, r *http.Request) {
	p := models.NewPereval()
	var err error
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		SendErr(w, r, http.StatusServiceUnavailable, err)
		return
	}
	if err := p.ValidateFields(); err != nil {
		SendErr(w, r, http.StatusBadRequest, err)
		return
	}
	if err := p.ValidateData(); err != nil {
		SendErr(w, r, http.StatusServiceUnavailable, err)
		return
	}
	var img [][]byte
	if img, err = GetImage(&p); err != nil {
		SendErr(w, r, http.StatusServiceUnavailable, err)
		return
	}
	id := make(map[string]string)
	if id, err = dbcontroller.AddImage(img, &p); err != nil {
		SendErr(w, r, http.StatusServiceUnavailable, err)
		return
	}
	imgMap := make(map[string][]int)
	//var ai []map[string][]int
	for key, title := range id {
		imgId, err := strconv.Atoi(key)
		if err != nil {
			SendErr(w, r, http.StatusServiceUnavailable, err)
			return
		}
		imgMap[title] = append(imgMap[title], imgId)
		//ai = append(ai, imgMap)
		/*		ai = append(ai, models.AddImages{
					ID:    key,
					Title: val,
				})
		*/
	}
	if id, err := dbcontroller.AddData(&p, &imgMap); err != nil {
		SendErr(w, r, http.StatusServiceUnavailable, err)
		return
	} else {
		SendResponse(w, id)
	}
}

func SendResponse(w http.ResponseWriter, data string) {
	resp := Response{
		Message: "OK",
		ID:      data,
	}
	jsonResp, _ := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResp)
}
