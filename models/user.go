package models

type User struct {
	ID     int      `json:"id"`
	Name   string   `json:"name"`
	Groups []*Group `gorm:"many2many:user_groups;"`
}

func NewUser(name string) *User {
	return &User{Name: name}
}
