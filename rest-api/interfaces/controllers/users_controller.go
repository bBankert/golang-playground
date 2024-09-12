package interfaces

import "github.com/gin-gonic/gin"

type IUsersController interface {
	CreateUser(context *gin.Context)
	Login(context *gin.Context)
}
