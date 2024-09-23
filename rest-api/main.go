package main

import (
	"fmt"

	"example.com/config"
)

func main() {

	err := config.LoadConfiguration()

	if err != nil {
		panic(fmt.Sprintf("Unable to load configuration, error: %v\n", err.Error()))
	}

	app, err := BuildServer()

	if err != nil {
		panic(fmt.Sprintf("Unable to build server, error: %v\n", err.Error()))
	}

	appConfig := config.AppConfiguration()

	httpPort, err := appConfig.HttpPort()

	if err != nil {
		panic(fmt.Sprintf("Unable to start server, err: %v\n", err.Error()))
	}

	err = app.Start(httpPort)

	if err != nil {
		panic(fmt.Sprintf("Unable to run server, error: %v\n", err.Error()))
	}
}
