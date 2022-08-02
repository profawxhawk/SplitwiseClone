package models

type Group struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Users []*User `gorm:"many2many:user_groups;"`
}

func NewGroup(name string) *Group {
	return &Group{Name: name}
}
