package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Configuration struct {
	HttpPort     string
	JwtSecretKey string
}

var Config Configuration

func LoadConfiguration() error {

	err := godotenv.Load()

	if err != nil {
		return err
	}

	Config = Configuration{
		HttpPort:     os.Getenv("HTTP_PORT"),
		JwtSecretKey: os.Getenv("TOKEN_SECRET"),
	}

	return nil
}
