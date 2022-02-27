package models

type Pereval struct {
	ID          string `json:"id"`
	BeautyTitle string `json:"beautyTitle"`
	Title       string `json:"title"`
	OtherTitles string `json:"other_titles"`
	Connect     string `json:"connect"`
	AddTime     string `json:"add_time"`
	User        struct {
		ID    string `json:"id"`
		Email string `json:"email"`
		Phone string `json:"phone"`
		Fam   string `json:"fam"`
		Name  string `json:"name"`
		Otc   string `json:"otc"`
	} `json:"user"`
	Coords struct {
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
