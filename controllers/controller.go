package controllers

import (
	"SplitwiseClone/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ControllerInterface interface {
	AddUserController(c *gin.Context)
	AddGroupController(c *gin.Context)
	AddUserToGroupController(c *gin.Context)
	AddExpenseToGroupController(c *gin.Context)
	GetAllExpensesController(c *gin.Context)
}

type Controller struct {
	services.EntityCreationServiceInterface
	services.TransactionServiceInterface
}

func NewController(AllServices *services.Services) *Controller {
	return &Controller{EntityCreationServiceInterface: AllServices.EntityCreationService, TransactionServiceInterface: AllServices.TransactionService}
}

func (controller *Controller) AddUserController(c *gin.Context) {
	var req GetNameRequest
	if !req.GetPostRequests(c) {
		return
	}
	userId, err := controller.CreatUserEntity(req.Name)
	if !SendServerErrorResponse(c, err) {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"userId": userId,
	})
}

func (controller *Controller) AddGroupController(c *gin.Context) {
	var req GetNameRequest
	if !req.GetPostRequests(c) {
		return
	}
	groupId, err := controller.CreatGroupEntity(req.Name)
	if !SendServerErrorResponse(c, err) {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"groupId": groupId,
	})
}

func (controller *Controller) AddUserToGroupController(c *gin.Context) {
	var req AddUserToGroupRequest
	if !req.GetPostRequests(c) {
		return
	}
	err := controller.CreateUserGroupMapping(req.UserId, req.GroupId)
	if !SendServerErrorResponse(c, err) {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "User added to group successfully",
	})
}

func (controller *Controller) AddExpenseToGroupController(c *gin.Context) {
	var req AddExpenseToGroupRequest
	if !req.GetPostRequests(c) {
		return
	}
	err := controller.GenerateExpenseRecord(req.UserId, req.GroupId, req.ExpenseName, req.TotalAmount, req.ExpenseType, req.UserAmountMap)
	if !SendServerErrorResponse(c, err) {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Expense created successfully",
	})
}

func (controller *Controller) GetAllExpensesController(c *gin.Context) {
	userIdTemp, _ := c.GetQuery("userId")
	userId, err := strconv.Atoi(userIdTemp)
	if !SendServerErrorResponse(c, err) {
		return
	}
	transactionViewList, errController := controller.GetAllExpensesForUser(userId)
	if !SendServerErrorResponse(c, errController) {
		return
	}
	c.JSON(http.StatusOK, transactionViewList)
}
