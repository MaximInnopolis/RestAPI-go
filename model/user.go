package model

import "RestAPI-go/utils"

type User struct {
	Name      string         `json:"name"`
	Surname   string         `json:"surname"`
	Login     string         `json:"login"`
	Password  string         `json:"password"`
	BirthDate utils.DateTime `json:"birth_date"`
	Status    UserStatus     `json:"status"`
	Role      UserRole       `json:"role"`
}
