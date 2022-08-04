package services

import (
	"SplitwiseClone/repositories"
)

type Services struct {
	EntityCreationService *EntityCreationService
	TransactionService    *TransactionService
}

func InitServices(repos *repositories.Repositories) *Services {
	entityCreationService := NewEntityCreationService(repos.UserRepo, repos.GroupRepo)
	transactionService := NewTransactionService(repos.ExpenseRepo, repos.GroupRepo, repos.UserRepo)
	return &Services{EntityCreationService: entityCreationService, TransactionService: transactionService}
}
