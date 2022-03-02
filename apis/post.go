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
	//m := make(map[string]string)
	var m map[string]string
	if m, err = dbcontroller.AddImage(img, &p); err != nil {
		SendErr(w, r, http.StatusServiceUnavailable, err)
		return
	}
	var imgMap *map[string][]int
	if imgMap, err = imgData(m); err != nil {
		SendErr(w, r, http.StatusServiceUnavailable, err)
		return
	}

	if id, err := dbcontroller.AddPereval(&p, imgMap); err != nil {
		SendErr(w, r, http.StatusServiceUnavailable, err)
		return
	} else {
		SendResponse(w, id)
	}
}

func imgData(m map[string]string) (*map[string][]int, error) {
	imgMap := make(map[string][]int)
	var err error
	for key, title := range m {
		imgId, err := strconv.Atoi(key)
		if err != nil {
			return nil, err
		}
		imgMap[title] = append(imgMap[title], imgId)
	}
	return &imgMap, err
}
