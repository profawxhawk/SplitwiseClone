package repositories

import (
	"SplitwiseClone/models"
	"gorm.io/gorm"
	"log"
	"time"
)

type UserRepoInterface interface {
	CreateUser(name string) (*models.User, error)
	CreateUserGroupEntry(userId int, groupId int) bool
	GetAllUserExpenses(userId int) ([]GetAllUserExpenseDTO, error)
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
	err := userRepo.db.Model(&models.User{ID: userId}).
		Association("Groups").
		Append(&models.Group{ID: groupId})
	if err != nil {
		return false
	}
	return true
}

type GetAllUserExpenseDTO struct {
	UserId      int       `json:"user_id"`
	GroupName   string    `json:"group_name"`
	ExpenseName string    `json:"expense_name"`
	PaidAmount  float32   `json:"paid_amount"`
	OwedSum     float32   `json:"owed_amount"`
	ShouldPay   bool      `json:"should_pay"`
	CreatedAt   time.Time `json:"date"`
}

func (userRepo *UserRepo) GetAllUserExpenses(userId int) ([]GetAllUserExpenseDTO, error) {
	var responseDTOList []GetAllUserExpenseDTO
	err := userRepo.db.Raw("? UNION ?",
		userRepo.db.Model(&models.UserExpense{}).Select("False as should_pay, user_expenses.paid_user_id as user_id,groups.name as group_name, expenses.name as expense_name,user_expenses.paid_amount, sum(user_expenses.owed_amount) as owed_sum,user_expenses.created_at as date").
			Joins("inner join expenses on user_expenses.expense_id = expenses.id").
			Joins("inner join groups on expenses.group_id = groups.id").
			Group("user_expenses.expense_id").Group("user_expenses.paid_user_id").Group("groups.name").
			Group("expenses.name").Group("user_expenses.paid_amount").
			Group("user_expenses.created_at").Group("user_expenses.pending").
			Having("(user_expenses.paid_user_id = ?) and user_expenses.pending=true", userId),
		userRepo.db.Model(&models.UserExpense{}).Select("True as should_pay, user_expenses.owed_user_id as user_id,groups.name as group_name, expenses.name as expense_name,user_expenses.paid_amount, sum(user_expenses.owed_amount) as owed_sum,user_expenses.created_at as date").
			Joins("inner join expenses on user_expenses.expense_id = expenses.id").
			Joins("inner join groups on expenses.group_id = groups.id").
			Group("user_expenses.expense_id").Group("groups.name").
			Group("expenses.name").Group("user_expenses.owed_user_id").Group("user_expenses.paid_amount").
			Group("user_expenses.created_at").Group("user_expenses.pending").
			Having("(user_expenses.owed_user_id = ?) and user_expenses.pending=true", userId),
	).Scan(&responseDTOList).Error

	return responseDTOList, err
}
