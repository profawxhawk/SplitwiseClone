package controllers

import (
	"SplitwiseClone/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ControllerInterface interface {
	AddUserController(c *gin.Context)
	AddGroupController(c *gin.Context)
}

type Controller struct {
	services.EntityCreationServiceInterface
}

func NewController(AllServices *services.Services) *Controller {
	return &Controller{EntityCreationServiceInterface: AllServices.EntityCreationService}
}

func (controller *Controller) AddUserController(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		c.AbortWithStatusJSON(400, gin.H{"message": "name parameter not found"})
		return
	}
	userId, err := controller.CreatUserEntity(name)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": "Server error, please try again"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"userId": userId,
	})
}

func (controller *Controller) AddGroupController(c *gin.Context) {
	name := c.Query("name")
	if name == "" {
		c.AbortWithStatusJSON(400, gin.H{"message": "name parameter not found"})
		return
	}
	groupId, err := controller.CreatGroupEntity(name)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": "Server error, please try again"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"groupId": groupId,
	})
}
