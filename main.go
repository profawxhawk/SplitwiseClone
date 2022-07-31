package main

import (
	"SplitwiseClone/api"
	"SplitwiseClone/configs"
	"SplitwiseClone/controllers"
	"SplitwiseClone/databases"
	"SplitwiseClone/repositories"
	"SplitwiseClone/services"
)

func main() {
	configs.LoadEnv()
	dbConn := databases.InitDB()
	repository := repositories.InitRepositories(dbConn)
	service := services.InitServices(repository)
	controller := controllers.NewController(service)
	router := api.NewRouter(controller)
	router.StartRouter()
}
