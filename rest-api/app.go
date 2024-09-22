package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	interfaces "example.com/interfaces/controllers"
	"example.com/routes"
)

type App struct {
	server       *gin.Engine
	httpHandlers *HTTPHandlers
}

func (app App) Start(port string) error {

	app.InitializeRoutes(*app.httpHandlers)

	err := app.server.Run(fmt.Sprintf(":%v", port))

	if err != nil {
		return err
	}

	return nil
}

func (app App) InitializeRoutes(httpHandlers HTTPHandlers) {
	routes.RegisterEventRoutes(app.server, app.httpHandlers.eventsController)
	routes.RegisterUserRoutes(app.server, app.httpHandlers.usersController)
	routes.RegisterRegistrationRoutes(app.server, app.httpHandlers.registrationsController)
}

func NewApp(httpServer *gin.Engine, httpHandlers *HTTPHandlers) *App {
	return &App{
		server:       httpServer,
		httpHandlers: httpHandlers,
	}
}

type HTTPHandlers struct {
	eventsController        interfaces.IEventsController
	usersController         interfaces.IUsersController
	registrationsController interfaces.IRegistrationsController
}

func NewHTTPHandlers(eventsController interfaces.IEventsController, usersController interfaces.IUsersController, registrationsConroller interfaces.IRegistrationsController) *HTTPHandlers {
	return &HTTPHandlers{
		eventsController:        eventsController,
		usersController:         usersController,
		registrationsController: registrationsConroller,
	}
}
