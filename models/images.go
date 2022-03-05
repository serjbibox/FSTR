package models

type Images struct {
	URL   string `json:"url" example:"https://img.icons8.com/ios/2x/roller-skating.png"`
	Title string `json:"title" example:"Спуск. Фото №99"`
}

type ImageLoaded struct {
	Title string `json:"title" example:"Спуск. Фото №99"`
	Blob  string `json:"blob"` //Строку удобнее тестировать
	//Blob  []byte `json:"blob"`
}

type ImagesMap struct {
	Title string `json:"title" example:"Спуск. Фото №99"`
	ID    []byte `json:"id"`
}

func NewImageLoaded() *ImageLoaded {
	return &ImageLoaded{}
}
