package models

import "github.com/google/uuid"

type Key struct {
	ID     int64  `json:"id"`
	Key    string `json:"key"`
	UserID int    `json:"user_id"`
}

func GenerateKey(id int) Key {

	return Key{
		Key:    uuid.New().String(),
		UserID: id,
	}
}
