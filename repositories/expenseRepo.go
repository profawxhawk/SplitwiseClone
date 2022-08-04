package repositories

import (
	"SplitwiseClone/models"
	"gorm.io/gorm"
	"log"
)

type ExpenseRepoInterface interface {
	CreateExpense([]models.UserExpense, *models.Expense) error
}

type ExpenseRepo struct {
	db *gorm.DB
}

func NewExpenseRepo(db *gorm.DB) *ExpenseRepo {
	err := db.AutoMigrate(&models.Group{}, &models.Expense{}, &models.UserExpense{})
	if err != nil {
		log.Fatal("Error auto migrating user repo")
	}
	return &ExpenseRepo{db: db}
}

func (expenseRepo *ExpenseRepo) CreateExpense(userExpense []models.UserExpense, expense *models.Expense) error {

	txn := expenseRepo.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			txn.Rollback()
		}
	}()
	if err := txn.Error; err != nil {
		return err
	}
	if err := txn.Create(expense).Error; err != nil {
		txn.Rollback()
		return err
	}

	for index := range userExpense {
		userExpenseElement := &userExpense[index]
		userExpenseElement.ExpenseId = expense.ID
	}

	if err := txn.Create(&userExpense).Error; err != nil {
		txn.Rollback()
		return err
	}

	return txn.Commit().Error

}
