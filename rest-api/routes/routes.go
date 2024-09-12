package routes

import (
	interfaces "example.com/interfaces/controllers"
	"example.com/middlewares"
	"github.com/gin-gonic/gin"
)

func NewHttpServer() *gin.Engine {
	return gin.Default()
}

func RegisterEventRoutes(server *gin.Engine, eventsController interfaces.IEventsController) {
	unauthenticatedEventEndpoints := server.Group("/events")
	{
		unauthenticatedEventEndpoints.GET("", eventsController.GetEvents)

		unauthenticatedEventEndpoints.GET(":id", eventsController.GetEventById)
	}

	authtenticatedEventEndpoints := server.Group("/events")
	{
		authtenticatedEventEndpoints.Use(middlewares.Authenticate)
		authtenticatedEventEndpoints.POST("", eventsController.AddEvent)
		authtenticatedEventEndpoints.PUT(":id", eventsController.UpdateEvent)
		authtenticatedEventEndpoints.DELETE(":id", eventsController.DeleteEvent)
	}
}

func RegisterUserRoutes(server *gin.Engine, userController interfaces.IUsersController) {
	server.POST("/signup", userController.CreateUser)
	server.POST("/login", userController.Login)
}

func RegisterRegistrationRoutes(server *gin.Engine, registrationsController interfaces.IRegistrationsController) {
	registationRoutes := server.Group("/events/:id")
	{
		registationRoutes.Use(middlewares.Authenticate)
		registationRoutes.POST("/register", registrationsController.RegisterForEvent)
		registationRoutes.DELETE("/unregister", registrationsController.CancelEventRegistration)
	}
}
