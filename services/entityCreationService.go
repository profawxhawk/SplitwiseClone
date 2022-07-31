package services

import "SplitwiseClone/repositories"

type EntityCreationServiceInterface interface {
	CreatUserEntity(name string) (int, error)
}

type EntityCreationService struct {
	repositories.UserRepoInterface
}

func NewEntityCreationService(userRepoInterface repositories.UserRepoInterface) *EntityCreationService {
	return &EntityCreationService{UserRepoInterface: userRepoInterface}
}

func (entityCreationService *EntityCreationService) CreatUserEntity(name string) (int, error) {
	user, err := entityCreationService.CreateUser(name)
	return user.ID, err
}
