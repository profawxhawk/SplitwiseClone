package models

type Group struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func NewGroup(name string) *Group {
	return &Group{Name: name}
}
