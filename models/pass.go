package models

type Pass struct {
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
	Status string   `json:"-"`
}

type PassAdded struct {
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
