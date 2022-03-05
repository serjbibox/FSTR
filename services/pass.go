package services

import (
	"errors"
	"net/http"

	"github.com/serjbibox/FSTR/models"
)

type passDAO interface {
	InsertTo(f *models.Flow, table string) *models.Flow
	Get(*models.Flow) *models.Flow
	Create(f *models.Flow, r *http.Request) *models.Flow
}

type PassService struct {
	dao passDAO
}

func New(dao passDAO) *PassService {
	return &PassService{dao}
}

func NewFlow() *models.Flow {
	return &models.Flow{
		Pass: &models.Pass{
			Type: "pass",
		},
		Warning: errors.New(""),
	}
}

func (s *PassService) Get(f *models.Flow) *models.Flow {
	return s.dao.Get(f)
}
func (s *PassService) Create(f *models.Flow, r *http.Request) *models.Flow {
	return s.dao.Create(f, r)
}
func (s *PassService) InsertTo(f *models.Flow, table string) *models.Flow {
	return s.dao.InsertTo(f, table)
}
