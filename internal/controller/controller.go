package controller

import (
	"github.com/HarukiFT/go-crud/internal/domain"
	"github.com/HarukiFT/go-crud/internal/service"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	engine  *gin.Engine
	usecase *service.Service

	userConroller domain.UserController
}

func NewController(e *gin.Engine, usecase *service.Service) *Controller {
	c := Controller{
		engine:  e,
		usecase: usecase,

		userConroller: NewUserController(e, &usecase.UserService),
	}

	return &c
}
