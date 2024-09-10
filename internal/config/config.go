package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	TCPPort     string
	KafkaBroker string
	HTTPPort    string
	RedisPort   string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	return &Config{
		TCPPort:     os.Getenv("TCP_PORT"),
		KafkaBroker: os.Getenv("KAFKA_BROKER"),
		HTTPPort:    os.Getenv("HTTP_PORT"),
		RedisPort:   os.Getenv("REDIS_PORT"),
	}, nil
}
