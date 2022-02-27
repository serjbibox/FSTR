package models

type User struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Fam   string `json:"fam"`
	Name  string `json:"name"`
	Otc   string `json:"otc"`
}
