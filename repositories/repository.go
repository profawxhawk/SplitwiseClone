package repositories

import "gorm.io/gorm"

type Repositories struct {
	UserRepo *UserRepo
}

func InitRepositories(db *gorm.DB) *Repositories {
	userRepo := NewUserRepo(db)
	return &Repositories{UserRepo: userRepo}
}
