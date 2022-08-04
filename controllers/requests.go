package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type requests interface {
	validateRequestData(c *gin.Context) bool
	GetPostRequests(c *gin.Context) bool
}
type GetNameRequest struct {
	Name string `json:"name"`
	requests
}

func (req *GetNameRequest) ValidateRequestData(c *gin.Context) bool {
	if req.Name == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "name parameter not found"})
		return false
	}
	return true
}
func (req *GetNameRequest) GetPostRequests(c *gin.Context) bool {
	if err := c.BindJSON(&req); err != nil {
		return SendServerErrorResponse(c, err)
	}
	return req.ValidateRequestData(c)
}

type AddUserToGroupRequest struct {
	UserId  int `json:"userId"`
	GroupId int `json:"groupId"`
	requests
}

func (req *AddUserToGroupRequest) ValidateRequestData(c *gin.Context) bool {
	if req.UserId == 0 || req.GroupId == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return false
	}
	return true
}

func (req *AddUserToGroupRequest) GetPostRequests(c *gin.Context) bool {
	if err := c.BindJSON(&req); err != nil {
		return SendServerErrorResponse(c, err)
	}
	return req.ValidateRequestData(c)
}

type AddExpenseToGroupRequest struct {
	UserId        int             `json:"userId"`
	GroupId       int             `json:"groupId"`
	ExpenseName   string          `json:"expenseName"`
	TotalAmount   float32         `json:"totalAmount"`
	ExpenseType   int             `json:"expenseType"` // go doesn't have enums by default, so using int
	UserAmountMap map[int]float32 `json:"userAmountMap"`
	requests
}

func (req *AddExpenseToGroupRequest) ValidateRequestData(c *gin.Context) bool {
	if req.UserId == 0 || req.GroupId == 0 || req.ExpenseName == "" || req.TotalAmount == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return false
	}
	if req.ExpenseType == 1 && req.UserAmountMap == nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Bad request"})
		return false
	}
	return true
}

func (req *AddExpenseToGroupRequest) GetPostRequests(c *gin.Context) bool {
	if err := c.BindJSON(&req); err != nil {
		return SendServerErrorResponse(c, err)
	}
	return req.ValidateRequestData(c)
}
