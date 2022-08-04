package repositories

import "gorm.io/gorm"

type Repositories struct {
	UserRepo    *UserRepo
	GroupRepo   *GroupRepo
	ExpenseRepo *ExpenseRepo
}

func InitRepositories(db *gorm.DB) *Repositories {
	userRepo := NewUserRepo(db)
	groupRepo := NewGroupRepo(db)
	expenseRepo := NewExpenseRepo(db)
	return &Repositories{UserRepo: userRepo, GroupRepo: groupRepo, ExpenseRepo: expenseRepo}
}
