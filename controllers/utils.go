package controllers

import "github.com/gin-gonic/gin"

func SendServerErrorResponse(c *gin.Context, err error) bool {
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return false
	}
	return true
}
