package models

import "errors"

type Pereval struct {
	ID          string `json:"id"`
	BeautyTitle string `json:"beautyTitle"`
	Title       string `json:"title"`
	OtherTitles string `json:"other_titles"`
	Connect     string `json:"connect"`
	AddTime     string `json:"add_time"`
	User        User   `json:"user"`
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
	Images []struct {
		URL   string `json:"url"`
		Title string `json:"title"`
	} `json:"images"`
}

func NewPereval() Pereval {
	return Pereval{
		Type: "pass",
	}
}

func (p *Pereval) Validate() error {
	if p.User.ID == "" {
		return errors.New("отсутствует ID пользователя")
	}
	if p.AddTime == "" {
		return errors.New("T0")
	}
	return nil
}
