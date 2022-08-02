package controllers

import (
	"SplitwiseClone/services"
	"fmt"
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

type GetNameRequest struct {
	Name string `json:"name"`
}

type AddUserToGroupRequest struct {
	UserId  int `json:"userId"`
	GroupId int `json:"groupId"`
}

func (controller *Controller) AddUserController(c *gin.Context) {
	var req GetNameRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(req)
	if req.Name == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "name parameter not found"})
		return
	}
	userId, err := controller.CreatUserEntity(req.Name)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": "Server error, please try again"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"userId": userId,
	})
}

func (controller *Controller) AddGroupController(c *gin.Context) {
	var req GetNameRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if req.Name == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "name parameter not found"})
		return
	}
	groupId, err := controller.CreatGroupEntity(req.Name)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": "Server error, please try again"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"groupId": groupId,
	})
}

func (controller *Controller) AddUserToGroupController(c *gin.Context) {
	var req AddUserToGroupRequest
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if req.UserId == 0 || req.GroupId == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Ids missing"})
		return
	}
	err := controller.CreateUserGroupMapping(req.UserId, req.GroupId)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": "Server error, please try again"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "User added to group successfully",
	})
}
