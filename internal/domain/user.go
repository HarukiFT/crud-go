package domain

import "github.com/gin-gonic/gin"

type UserController interface {
	NewUser(context *gin.Context)
}

type UserService interface {
	NewUser(name string, password string) bool
}
