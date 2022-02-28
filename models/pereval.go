package models

import (
	"errors"
	"fmt"
	"time"
)

type PerevalAdded struct {
	ID          string `json:"pereval_id" example:"125"`
	BeautyTitle string `json:"beautyTitle" example:"пер. "`
	Title       string `json:"title" example:"Туя-Ашуу"`
	OtherTitles string `json:"other_titles" example:"1"`
	Connect     string `json:"connect" example:" "`
	AddTime     string `json:"add_time" example:"2021-09-22 13:18:13"`
	Coords      struct {
		Latitude  string `json:"latitude" example:"45.3842"`
		Longitude string `json:"longitude" example:"7.1525"`
		Height    string `json:"height" example:"1200"`
	} `json:"coords"`
	Type  string `json:"type" example:"pass"`
	Level struct {
		Winter string `json:"winter" example:" "`
		Summer string `json:"summer" example:"1A"`
		Autumn string `json:"autumn" example:"1A"`
		Spring string `json:"spring" example:" "`
	} `json:"level"`
	User User `json:"user"`
}

func NewPerevalAdded(p *Pereval) PerevalAdded {
	return PerevalAdded{
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

type Pereval struct {
	ID          string `json:"pereval_id" example:"125"`
	BeautyTitle string `json:"beautyTitle" example:"пер. "`
	Title       string `json:"title" example:"Туя-Ашуу"`
	OtherTitles string `json:"other_titles" example:"1"`
	Connect     string `json:"connect" example:" "`
	AddTime     string `json:"add_time" example:"2021-09-22 13:18:13"`
	Coords      struct {
		Latitude  string `json:"latitude" example:"45.3842"`
		Longitude string `json:"longitude" example:"7.1525"`
		Height    string `json:"height" example:"1200"`
	} `json:"coords"`
	Type  string `json:"type" example:"pass"`
	Level struct {
		Winter string `json:"winter" example:" "`
		Summer string `json:"summer" example:"1A"`
		Autumn string `json:"autumn" example:"1A"`
		Spring string `json:"spring" example:" "`
	} `json:"level"`
	User   User     `json:"user"`
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
		return errors.New("отсутствует ID перевала")
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
