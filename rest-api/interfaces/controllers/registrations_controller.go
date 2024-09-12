package interfaces

import "github.com/gin-gonic/gin"

type IRegistrationsController interface {
	RegisterForEvent(context *gin.Context)
	CancelEventRegistration(context *gin.Context)
}
