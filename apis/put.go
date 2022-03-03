package apis

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/render"
	"github.com/serjbibox/FSTR/dbcontroller"
	"github.com/serjbibox/FSTR/models"
)

func UpdatePass(w http.ResponseWriter, r *http.Request) {
	if ctx, ok := r.Context().Value("pass").(*Context); !ok {
		err := errors.New("ошибка контекста GetStatus")
		SendErr(w, r, http.StatusServiceUnavailable, fmt.Errorf("%w", err))
		return
	} else {
		if ctx.Status.Status == "new" {
			doit(w, r, ctx.Pereval)
		}

		if err := render.Render(w, r, &PerevalResponse{Status: *ctx.Status}); err != nil {
			SendErr(w, r, http.StatusServiceUnavailable, fmt.Errorf("%w", err))
			return
		}
	}

}

func doit(w http.ResponseWriter, r *http.Request, p *models.Pereval) {
	var err error
	var p int
	log.Println(p)
	pnew := models.NewPereval()
	if err := json.NewDecoder(r.Body).Decode(&pnew); err != nil {
		SendErr(w, r, http.StatusServiceUnavailable, err)
		return
	}
	if err := pnew.ValidateFields(); err != nil {
		SendErr(w, r, http.StatusBadRequest, err)
		return
	}
	if err := pnew.ValidateData(); err != nil {
		SendErr(w, r, http.StatusServiceUnavailable, err)
		return
	}
	pnew.User = p.User
	var img [][]byte
	if img, err = GetImage(&pnew); err != nil {
		SendErr(w, r, http.StatusServiceUnavailable, err)
		return
	}
	var m map[string]string
	if m, err = dbcontroller.AddImage(img, &pnew); err != nil {
		SendErr(w, r, http.StatusServiceUnavailable, err)
		return
	}
	var imgMap *map[string][]int
	if imgMap, err = imgData(m); err != nil {
		SendErr(w, r, http.StatusServiceUnavailable, err)
		return
	}

	if id, err := dbcontroller.AddPereval(&pnew, imgMap); err != nil {
		SendErr(w, r, http.StatusServiceUnavailable, err)
		return
	} else {
		SendResponse(w, id)
	}
}
