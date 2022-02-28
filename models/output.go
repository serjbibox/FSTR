package models

// @Description Структура HTTP ответа:
// @Description если отправка успешна, дополнительно возвращается id вставленной записи.
type Response struct {
	Message string `json:"message" example:"OK"`
	ID      string `json:"id" example:"123"`
}

// @Description Структура HTTP ответа об ошибке
type ResponseErr struct {
	Message string `json:"message" example:"ошибка: описание ошибки"`
}

func NewResponse() Response {
	return Response{}
}
