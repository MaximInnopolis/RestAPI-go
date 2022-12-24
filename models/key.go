package models

type Key struct {
	ID     int64  `json:"id"`
	Key    string `json:"key"`
	UserID int    `json:"user_id"`
}

func GenerateKey(id int) Key {
	return Key{
		Key:    "string",
		UserID: id,
	}
}
