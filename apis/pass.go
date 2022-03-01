package apis

import "net/http"

var PassRows = []*Pass{
	{ID: "4", Title: "4"},
	{ID: "6", Title: "6"},
}

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
}

type User struct {
	ID    string `json:"id" example:"11234"`
	Email string `json:"email" example:"dd@aa.ru"`
	Phone string `json:"phone" example:"+744434555"`
	Fam   string `json:"fam" example:"Скворцов"`
	Name  string `json:"name" example:"Иван"`
	Otc   string `json:"otc" example:"Кожедубович"`
}

type Images struct {
	URL   string `json:"url" example:"http://..."`
	Title string `json:"title" example:"Спуск. Фото №99"`
}

func NewPassResponse(pass *Pass) *PassResponse {
	resp := &PassResponse{Pass: pass}
	return resp
}

type PassResponse struct {
	*Pass
	// We add an additional field to the response here.. such as this
	// elapsed computed property
	Elapsed int64 `json:"elapsed"`
}

func (rd *PassResponse) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing before a response is marshalled and sent across the wire
	rd.Elapsed = 10
	return nil
}
