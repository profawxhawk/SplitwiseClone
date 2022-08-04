package services

import (
	"SplitwiseClone/models"
	"SplitwiseClone/repositories"
	"errors"
)

type TransactionServiceInterface interface {
	GenerateExpenseRecord(userId int, groupId int, expenseName string, totalAmount float32, expenseType int, userAmountMap map[int]float32) error
	GetAllExpensesForUser(userId int) ([]TransactionView, error)
}

type TransactionService struct {
	repositories.GroupRepoInterface
	repositories.ExpenseRepoInterface
	repositories.UserRepoInterface
}

func NewTransactionService(expenseRepoInterface repositories.ExpenseRepoInterface, groupRepoInterface repositories.GroupRepoInterface, userRepoInterface repositories.UserRepoInterface) *TransactionService {
	return &TransactionService{ExpenseRepoInterface: expenseRepoInterface, GroupRepoInterface: groupRepoInterface, UserRepoInterface: userRepoInterface}
}
func (transactionService *TransactionService) GenerateExpenseRecord(userId int, groupId int, expenseName string, totalAmount float32, expenseType int, userAmountMap map[int]float32) error {
	expense := models.NewExpense(expenseName, groupId)
	var userExpenseList []models.UserExpense
	allUserIds, err := transactionService.GetAllUserIdsInGroupExcludingGivenUser(groupId, userId)
	if err != nil {
		return err
	}
	allUserIdSet := make(map[int]bool)
	for i := 0; i < len(allUserIds); i += 1 {
		allUserIdSet[allUserIds[i]] = true
	}
	if expenseType == 0 { //equal case

		if len(allUserIds) == 0 {
			return errors.New("group has only one user")
		}

		for _, owedUserId := range allUserIds {
			userExpenseList = append(userExpenseList,
				models.UserExpense{PaidUser: models.User{ID: userId},
					PaidAmount: totalAmount, OwedUser: models.User{ID: owedUserId}, OwedAmount: totalAmount / float32(len(allUserIds)), Pending: true})
		}
	} else if expenseType == 1 {
		var totalOwedUserAmount float32 = 0.0
		for owedUserId, owedUserAmount := range userAmountMap {
			userExpenseList = append(userExpenseList,
				models.UserExpense{PaidUser: models.User{ID: userId},
					PaidAmount: totalAmount, OwedUser: models.User{ID: owedUserId}, OwedAmount: owedUserAmount, Pending: true})
			totalOwedUserAmount += owedUserAmount
			if !allUserIdSet[owedUserId] {
				return errors.New("incorrect user id found in map")
			}
		}
		if totalOwedUserAmount != totalAmount {
			return errors.New("invalid amount split")
		}
	}
	errExpense := transactionService.CreateExpense(userExpenseList, expense)
	if errExpense != nil {
		return errExpense
	}
	return nil
}

type TransactionView struct {
	Date                   string
	Group                  string
	Expense                string
	TotalAmount            float32
	UserAmountContribution float32
}

func (transactionService *TransactionService) GetAllExpensesForUser(userId int) ([]TransactionView, error) {
	responseDTOList, err := transactionService.GetAllUserExpenses(userId)
	if err != nil {
		return nil, err
	}
	var transactionViewList []TransactionView
	dateFormat := "2000-01-13"
	for _, responseDTO := range responseDTOList {
		formattedDate := responseDTO.CreatedAt.Format(dateFormat)
		transactionView := TransactionView{Date: formattedDate, Group: responseDTO.GroupName, Expense: responseDTO.ExpenseName, TotalAmount: responseDTO.PaidAmount}
		if !responseDTO.ShouldPay {
			transactionView.UserAmountContribution = responseDTO.OwedSum
		} else {
			transactionView.UserAmountContribution = -responseDTO.OwedSum
		}
		transactionViewList = append(transactionViewList, transactionView)
	}
	return transactionViewList, nil
}
