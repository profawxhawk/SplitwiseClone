package repositories

import (
	"SplitwiseClone/models"
	"gorm.io/gorm"
	"log"
)

type GroupRepoInterface interface {
	CreateGroup(name string) (*models.Group, error)
	GetAllUserIdsInGroupExcludingGivenUser(groupId int, userId int) ([]int, error)
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

func (groupRepo *GroupRepo) GetAllUserIdsInGroupExcludingGivenUser(groupId int, userId int) ([]int, error) {
	userIdList := make([]int, 0)
	err := groupRepo.db.Model(&models.User{}).Preload("Groups", "id != ?", groupId).
		Where("id != ?", userId).
		Select("id").Find(&userIdList).Error
	return userIdList, err
}
