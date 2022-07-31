package repositories

import "gorm.io/gorm"

type Repositories struct {
	userRepo *UserRepo
}

func InitRepositories(db *gorm.DB) *Repositories {
	userRepo := NewUserRepo(db)
	return &Repositories{userRepo: userRepo}
}
