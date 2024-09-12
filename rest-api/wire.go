//go:build wireinject
// +build wireinject

package main

import (
	"example.com/config"
	"example.com/controllers"
	controllerInterfaces "example.com/interfaces/controllers"
	repositoryInterfaces "example.com/interfaces/repositories"
	serviceInterfaces "example.com/interfaces/services"
	"example.com/repositories"
	"example.com/routes"
	"example.com/services"

	"github.com/google/wire"
)

func BuildServer() (*App, error) {
	wire.Build(
		config.InitializeDatabase,
		//repository registration
		repositories.NewEventRepository,
		wire.Bind(new(repositoryInterfaces.IEventRepository), new(*repositories.EventRepository)),
		repositories.NewRegistrationRepository,
		wire.Bind(new(repositoryInterfaces.IRegistrationRepository), new(*repositories.RegistrationRepository)),
		repositories.NewUserRepository,
		wire.Bind(new(repositoryInterfaces.IUserRepository), new(*repositories.UserRepository)),
		//service registration
		services.NewEventService,
		wire.Bind(new(serviceInterfaces.IEventService), new(*services.EventService)),
		services.NewUserService,
		wire.Bind(new(serviceInterfaces.IUserService), new(*services.UserService)),
		services.NewRegistrationService,
		wire.Bind(new(serviceInterfaces.IRegistrationService), new(*services.RegistrationService)),
		//controller registration
		controllers.NewEventsController,
		wire.Bind(new(controllerInterfaces.IEventsController), new(*controllers.EventsController)),
		controllers.NewUsersController,
		wire.Bind(new(controllerInterfaces.IUsersController), new(*controllers.UsersController)),
		controllers.NewRegistrationsController,
		wire.Bind(new(controllerInterfaces.IRegistrationsController), new(*controllers.RegistrationsController)),
		routes.NewHttpServer,
		NewHTTPHandlers,
		NewApp,
	)

	return &App{}, nil
}
