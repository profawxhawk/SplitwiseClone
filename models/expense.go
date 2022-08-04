package models

import "time"

type Expense struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	GroupID int
	Group   Group `gorm:"foreignKey:GroupID"`
}

func NewExpense(name string, groupId int) *Expense {
	return &Expense{Name: name, Group: Group{ID: groupId}}
}

type UserExpense struct {
	ID         int `json:"id"`
	ExpenseId  int
	Expense    Expense `gorm:"foreignKey:ExpenseId"`
	PaidUserId int
	PaidUser   User    `gorm:"foreignKey:PaidUserId"`
	PaidAmount float32 `json:"paidAmount"`
	OwedUserId int
	OwedUser   User      `gorm:"foreignKey:OwedUserId"`
	OwedAmount float32   `json:"owedAmount"`
	CreatedAt  time.Time `json:"createdAt"`
	Pending    bool
}
