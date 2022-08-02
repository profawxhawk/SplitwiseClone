package repositories

import "gorm.io/gorm"

type Repositories struct {
	UserRepo  *UserRepo
	GroupRepo *GroupRepo
}

func InitRepositories(db *gorm.DB) *Repositories {
	userRepo := NewUserRepo(db)
	groupRepo := NewGroupRepo(db)
	return &Repositories{UserRepo: userRepo, GroupRepo: groupRepo}
}
