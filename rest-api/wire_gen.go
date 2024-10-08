// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"example.com/config"
	"example.com/controllers"
	"example.com/lib"
	"example.com/repositories"
	"example.com/routes"
	"example.com/services"
)

// Injectors from wire.go:

func BuildServer() (*App, error) {
	engine := routes.NewHttpServer()
	db := config.InitializeDatabase()
	eventRepository := repositories.NewEventRepository(db)
	eventService := services.NewEventService(eventRepository)
	eventsController := controllers.NewEventsController(eventService)
	userRepository := repositories.NewUserRepository(db)
	hasher := lib.NewHasher()
	userService := services.NewUserService(userRepository, hasher)
	jwtAuthorizer := lib.NewJwtAuthorizer()
	usersController := controllers.NewUsersController(userService, jwtAuthorizer)
	registrationRepository := repositories.NewRegistrationRepository(db)
	registrationService := services.NewRegistrationService(registrationRepository, eventRepository)
	registrationsController := controllers.NewRegistrationsController(registrationService)
	httpHandlers := NewHTTPHandlers(eventsController, usersController, registrationsController)
	app := NewApp(engine, httpHandlers)
	return app, nil
}
