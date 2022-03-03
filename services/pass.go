package services

import (
	"net/http"

	"github.com/serjbibox/FSTR/models"
)

type passDAO interface {
	Get(id string) (*models.Pass, error)
	Create(r *http.Request) (*models.Pass, error)
	Insert(p *models.Pass, imgMap *map[string][]int) (id string, err error)
	InsertImage(p *models.Pass, img [][]byte) (m map[string]string, err error)
	ValidateFields(*models.Pass) error
	ValidateData(*models.Pass) error
}

type PassService struct {
	dao passDAO
}

func New(dao passDAO) *PassService {
	return &PassService{dao}
}

func (s *PassService) Get(id string) (*models.Pass, error) {
	return s.dao.Get(id)
}

func (s *PassService) Create(r *http.Request) (*models.Pass, error) {
	return s.dao.Create(r)
}

func (s *PassService) ValidateFields(p *models.Pass) error {
	return s.dao.ValidateFields(p)
}

func (s *PassService) ValidateData(p *models.Pass) error {
	return s.dao.ValidateData(p)
}
func (s *PassService) Insert(p *models.Pass, imgMap *map[string][]int) (id string, err error) {
	return s.dao.Insert(p, imgMap)
}
func (s *PassService) InsertImage(p *models.Pass, img [][]byte) (m map[string]string, err error) {
	return s.dao.InsertImage(p, img)
}
