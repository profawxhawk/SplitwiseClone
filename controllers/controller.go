package controllers

import (
	"SplitwiseClone/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ControllerInterface interface {
	AddUserController(c *gin.Context)
	AddGroupController(c *gin.Context)
	AddUserToGroupController(c *gin.Context)
}

type Controller struct {
	services.EntityCreationServiceInterface
}

func NewController(AllServices *services.Services) *Controller {
	return &Controller{EntityCreationServiceInterface: AllServices.EntityCreationService}
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
	err := controller.CreateUserGroupMapping(req.UserId, req.GroupId)
	if !SendServerErrorResponse(c, err) {
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "User added to group successfully",
	})
}
