package models

type Images struct {
	URL   string `json:"url" example:"https://avatars.mds.yandex.net/i?id=a467876d3e1b1f0a84050103a206cf81-5858922-images-thumbs&n=13"`
	Title string `json:"title" example:"Спуск. Фото №99"`
}

type AddImages struct {
	ID    string `json:"id" example:"2"`
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
