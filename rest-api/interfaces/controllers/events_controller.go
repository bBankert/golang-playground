package interfaces

import "github.com/gin-gonic/gin"

type IEventsController interface {
	GetEvents(context *gin.Context)
	AddEvent(context *gin.Context)
	GetEventById(context *gin.Context)
	UpdateEvent(context *gin.Context)
	DeleteEvent(context *gin.Context)
}
