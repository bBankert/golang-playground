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

	err = app.Start(config.Config.HttpPort)

	if err != nil {
		panic(fmt.Sprintf("Unable to run server, error: %v\n", err.Error()))
	}
}
