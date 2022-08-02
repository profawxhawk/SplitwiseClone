package services

import (
	"SplitwiseClone/repositories"
)

type Services struct {
	EntityCreationService *EntityCreationService
}

func InitServices(repos *repositories.Repositories) *Services {
	entityCreationService := NewEntityCreationService(repos.UserRepo, repos.GroupRepo)
	return &Services{EntityCreationService: entityCreationService}
}
