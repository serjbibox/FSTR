package daos

import (
	"github.com/serjbibox/FSTR/models"
)

type PassDAO struct{}

func NewPassDAO() *PassDAO {
	return &PassDAO{}
}

func NewPassAdded(p *models.Pass) models.PassAdded {
	return models.PassAdded{
		ID:          p.ID,
		BeautyTitle: p.BeautyTitle,
		Title:       p.Title,
		OtherTitles: p.OtherTitles,
		Connect:     p.Connect,
		AddTime:     p.AddTime,
		Coords:      p.Coords,
		Type:        p.Type,
		Level:       p.Level,
		User:        p.User,
	}
}
