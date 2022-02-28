package models

type User struct {
	ID    string `json:"id" example:"11234"`
	Email string `json:"email" example:"dd@aa.ru"`
	Phone string `json:"phone" example:"+744434555"`
	Fam   string `json:"fam" example:"Скворцов"`
	Name  string `json:"name" example:"Иван"`
	Otc   string `json:"otc" example:"Кожедубович"`
}
