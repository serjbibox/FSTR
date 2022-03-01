package apis

import (
	"log"
	"net/http"

	"github.com/serjbibox/FSTR/models"

	"github.com/go-chi/render"
)

func GetPass(w http.ResponseWriter, r *http.Request) {
	if p, ok := r.Context().Value("pass").(*models.Pereval); !ok {
		//log.Println("не получилось", p.AddTime)
		log.Println("не получилось", r.Context().Value("pass"))
	} else {
		if err := render.Render(w, r, NewPerevalResponse(p)); err != nil {
			render.Render(w, r, ErrRender(err))
			return
		}
	}

	/*if err := render.Render(w, r, NewPerevalResponse(pointPereval)); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}*/
}

func NewPerevalResponse(p *models.Pereval) *PerevalResponse {
	resp := &PerevalResponse{Pereval: p}
	return resp
}

type PerevalResponse struct {
	*models.Pereval
	// We add an additional field to the response here.. such as this
	// elapsed computed property
	Elapsed int64 `json:"elapsed"`
}

func (rd *PerevalResponse) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing before a response is marshalled and sent across the wire
	rd.Elapsed = 10
	return nil
}
