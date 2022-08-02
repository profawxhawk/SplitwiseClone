package repositories

import (
	"SplitwiseClone/models"
	"gorm.io/gorm"
	"log"
)

type UserRepoInterface interface {
	CreateUser(name string) (*models.User, error)
	CreateUserGroupEntry(userId int, groupId int) bool
}

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	err := db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Error auto migrating user repo")
	}
	return &UserRepo{db: db}
}

func (userRepo *UserRepo) CreateUser(name string) (*models.User, error) {
	user := models.NewUser(name)
	result := userRepo.db.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (userRepo *UserRepo) CreateUserGroupEntry(userId int, groupId int) bool {
	result := userRepo.db.Model(&models.User{ID: userId}).Association("Groups").Append(&models.Group{ID: groupId})
	if result.Error != nil {
		return false
	}
	return true
}
