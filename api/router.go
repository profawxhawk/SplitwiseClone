package api

import (
	"SplitwiseClone/controllers"
	"github.com/gin-gonic/gin"
	"log"
)

type Router struct {
	ginRouter *gin.Engine
	controllers.ControllerInterface
}

func NewRouter(controllerInterface controllers.ControllerInterface) *Router {
	return &Router{ginRouter: gin.Default(), ControllerInterface: controllerInterface}
}

func (router *Router) StartRouter() {

	router.ginRouter.GET("/api/users/add", router.AddUserController)

	err := router.ginRouter.Run(":8080")
	if err != nil {
		log.Fatal("Error running router instance")
	}
}
