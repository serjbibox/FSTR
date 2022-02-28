package models

type Images struct {
	URL   string `json:"url" example:"http://..."`
	Title string `json:"title" example:"Спуск. Фото №99"`
}

/*
{
	"sedlo": [2,3],
	"Nord": [1],
	"West": null,
	"South": [4,5],
	"East": [6]
}
*/
