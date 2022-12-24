package models

import (
	"RestAPI-go/utils"
)

type User struct {
	ID        int            `json:"id"`
	Name      string         `json:"name"`
	Surname   string         `json:"surname"`
	Login     string         `json:"login"`
	Password  string         `json:"password"`
	BirthDate utils.DateTime `json:"birth_date"`
	Status    UserStatus     `json:"status"`
	Role      UserRole       `json:"role"`
}
