package model

type User struct {
	Name      string `json:"name"`
	Surname   string `json:"surname"`
	Login     string `json:"login"`
	Password  string `json:"password"`
	BirthDate string `json:"birth_date"`
}
