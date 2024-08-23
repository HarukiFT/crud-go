package main

import (
	"github.com/HarukiFT/go-crud/config"
	"github.com/HarukiFT/go-crud/internal/controller"
	"github.com/HarukiFT/go-crud/internal/service"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	db := config.GetDB()

	router := gin.Default()

	appService := service.NewService(db)
	controller.NewController(router, appService)

	router.Run(":8000")
}
