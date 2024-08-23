package controller

import (
	"net/http"

	"github.com/HarukiFT/go-crud/internal/domain"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	engine      *gin.Engine
	userService *domain.UserService
}

func (c *UserController) NewUser(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func (c *UserController) initRoutes() {
	userGroup := c.engine.Group("/user")
	userGroup.POST("/register", c.NewUser)
}

func NewUserController(e *gin.Engine, userService *domain.UserService) *UserController {
	c := UserController{engine: e, userService: userService}

	c.initRoutes()

	return &c
}
