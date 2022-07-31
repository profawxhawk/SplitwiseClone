package models

type User struct {
	ID   int    `json:"id"`
	name string `json:"name"`
}

func GetUser() User {
	var post User
	return post
}
