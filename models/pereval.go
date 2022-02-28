package models

import (
	"errors"
	"fmt"
	"time"
)

type Pereval struct {
	ID          string `json:"id"`
	BeautyTitle string `json:"beautyTitle"`
	Title       string `json:"title"`
	OtherTitles string `json:"other_titles"`
	Connect     string `json:"connect"`
	AddTime     string `json:"add_time"`
	ParsedTime  time.Time
	User        User `json:"user"`
	Coords      struct {
		Latitude  string `json:"latitude"`
		Longitude string `json:"longitude"`
		Height    string `json:"height"`
	} `json:"coords"`
	Type  string `json:"type"`
	Level struct {
		Winter string `json:"winter"`
		Summer string `json:"summer"`
		Autumn string `json:"autumn"`
		Spring string `json:"spring"`
	} `json:"level"`
	Images []Images `json:"images"`
}

func NewPereval() Pereval {
	return Pereval{
		Type: "pass",
	}
}

func (p *Pereval) ValidateFields() error {
	switch {
	case p.ID == "":
		return errors.New("отсутствует ID записи")
	case p.User.ID == "":
		return errors.New("отсутствует ID пользователя")
	case p.Coords.Height == "":
		return errors.New("отсутствует координата: coords.Height")
	case p.Coords.Latitude == "":
		return errors.New("отсутствует координата: coords.Latitude")
	case p.Coords.Longitude == "":
		return errors.New("отсутствует координата: coords.Longitude")
	case p.User.Name+p.User.Fam+p.User.Otc == "":
		return errors.New("отсутствует имя пользователя")
	case p.BeautyTitle+p.Title+p.OtherTitles == "":
		return errors.New("отсутствует название объекта")
	case len(p.Images) == 0:
		return errors.New("отсутствуют изображения")
	case len(p.Images) != 0:
		for idx, elem := range p.Images {
			if elem.URL == "" {
				s := fmt.Sprintf("отсутствует URL изображения: №%d, описание: %s", idx+1, elem.Title)
				return errors.New(s)
			}
		}
	}

	return nil
}

func (p *Pereval) ValidateData() error {
	if p.AddTime == "" {
		p.AddTime = "NOW()"
	} else if t, err := time.Parse("2006-01-02 15:04:05", p.AddTime); err != nil {
		return fmt.Errorf("%w", err)
	} else {
		p.AddTime = t.Format("2006-01-02 15:04:05")
	}
	return nil
}

/*
координаты объекта;
имя пользователя (ФИО строкой);
почту;
телефон пользователя;
название объекта;
несколько фотографий // кладутся в таблицу pereval_images (id, date_added, blob).
*/
