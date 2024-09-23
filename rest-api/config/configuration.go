package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type Configuration struct {
	httpPort     string
	jwtSecretKey string
}

var config Configuration

func LoadConfiguration() error {

	err := godotenv.Load()

	if err != nil {
		return err
	}

	config = Configuration{
		httpPort:     os.Getenv("HTTP_PORT"),
		jwtSecretKey: os.Getenv("TOKEN_SECRET"),
	}

	return nil
}

func (config Configuration) HttpPort() (string, error) {
	if config.httpPort == "" {
		return "", errors.New("missing http port configuration")
	}

	return config.httpPort, nil
}

func (config Configuration) JwtSecretKey() (string, error) {
	if config.jwtSecretKey == "" {
		return "", errors.New("missing auth secret configuration")
	}

	return config.jwtSecretKey, nil
}

func AppConfiguration() Configuration {
	return config
}
