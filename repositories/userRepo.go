package repositories

import (
	"SplitwiseClone/models"
	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db: db}
}
func (userRepo *UserRepo) CreateUser(name string) (*models.User, error) {
	user := &models.User{}
	result := userRepo.db.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}
