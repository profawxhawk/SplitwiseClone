package repositories

import (
	"SplitwiseClone/models"
	"gorm.io/gorm"
	"log"
)

type GroupRepoInterface interface {
	CreateGroup(name string) (*models.Group, error)
}

type GroupRepo struct {
	db *gorm.DB
}

func NewGroupRepo(db *gorm.DB) *GroupRepo {
	err := db.AutoMigrate(&models.Group{})
	if err != nil {
		log.Fatal("Error auto migrating user repo")
	}
	return &GroupRepo{db: db}
}
func (groupRepo *GroupRepo) CreateGroup(name string) (*models.Group, error) {
	group := models.NewGroup(name)
	result := groupRepo.db.Create(&group)
	if result.Error != nil {
		return nil, result.Error
	}
	return group, nil
}
