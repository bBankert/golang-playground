package routes

import (
	"database/sql"

	"example.com/controllers"
	interfaces "example.com/interfaces/controllers"
	"example.com/middlewares"
	"example.com/repositories"
	"example.com/services"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine, db *sql.DB) {

	eventsController, registrationsController, usersController := setupDependencies(db)

	registerEventRoutes(server, eventsController)
	registerUserRoutes(server, usersController)
	registerRegistrationRoutes(server, registrationsController)
}

func setupDependencies(db *sql.DB) (
	interfaces.IEventsController,
	interfaces.IRegistrationsController,
	interfaces.IUsersController,
) {
	//TODO: See if we can move these to a dependency injection section
	//repositories
	registrationRepository := repositories.NewRegistrationRepository(db)
	userRepository := repositories.NewUserRepository(db)
	eventRepository := repositories.NewEventRepository(db)

	//services
	eventService := services.NewEventService(eventRepository)
	userService := services.NewUserService(userRepository)
	registrationService := services.NewRegistrationService(
		registrationRepository,
		eventRepository)

	//controllers
	eventsController := controllers.NewEventsController(eventService)
	userController := controllers.NewUsersController(userService)
	registrationController := controllers.NewRegistrationsController(registrationService)

	return eventsController, registrationController, userController

}

func registerEventRoutes(server *gin.Engine, eventsController interfaces.IEventsController) {

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

func registerUserRoutes(server *gin.Engine, userController interfaces.IUsersController) {

	server.POST("/signup", userController.CreateUser)
	server.POST("/login", userController.Login)
}

func registerRegistrationRoutes(server *gin.Engine, registrationsController interfaces.IRegistrationsController) {

	registationRoutes := server.Group("/events/:id")
	{
		registationRoutes.Use(middlewares.Authenticate)
		registationRoutes.POST("/register", registrationsController.RegisterForEvent)
		registationRoutes.DELETE("/unregister", registrationsController.CancelEventRegistration)
	}
}
