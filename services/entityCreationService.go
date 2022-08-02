package services

import (
	"SplitwiseClone/repositories"
	"errors"
)

type EntityCreationServiceInterface interface {
	CreatUserEntity(name string) (int, error)
	CreatGroupEntity(name string) (int, error)
	CreateUserGroupMapping(userId int, groupId int) error
}

type EntityCreationService struct {
	repositories.UserRepoInterface
	repositories.GroupRepoInterface
}

func NewEntityCreationService(userRepoInterface repositories.UserRepoInterface, groupRepoInterface repositories.GroupRepoInterface) *EntityCreationService {
	return &EntityCreationService{UserRepoInterface: userRepoInterface, GroupRepoInterface: groupRepoInterface}
}

func (entityCreationService *EntityCreationService) CreatUserEntity(name string) (int, error) {
	user, err := entityCreationService.CreateUser(name)
	return user.ID, err
}

func (entityCreationService *EntityCreationService) CreatGroupEntity(name string) (int, error) {
	group, err := entityCreationService.CreateGroup(name)
	return group.ID, err
}

func (entityCreationService *EntityCreationService) CreateUserGroupMapping(userId int, groupId int) error {
	status := entityCreationService.CreateUserGroupEntry(userId, groupId)
	if !status {
		return errors.New("failed to create user group mapping")
	}
	return nil
}
