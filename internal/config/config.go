package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	TCPPort     string
	KafkaBroker string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load("../../.env")
	if err != nil {
		return nil, err
	}

	return &Config{
		TCPPort:     os.Getenv("TCP_PORT"),
		KafkaBroker: os.Getenv("KAFKA_BROKER"),
	}, nil
}
