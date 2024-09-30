package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	TCPPort          string
	KafkaBroker      string
	HTTPPort         string
	RedisPort        string
	ContainerPort    string
	PostgresPort     string
	PostgresUser     string
	PostgresPassword string
	PostgresName     string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	return &Config{
		TCPPort:          os.Getenv("TCP_PORT"),
		KafkaBroker:      os.Getenv("KAFKA_BROKER"),
		HTTPPort:         os.Getenv("HTTP_PORT"),
		RedisPort:        os.Getenv("REDIS_PORT"),
		ContainerPort:    os.Getenv("CONTAINER_PORT"),
		PostgresPort:     os.Getenv("POSTGRES_PORT"),
		PostgresUser:     os.Getenv("POSTGRES_USER"),
		PostgresPassword: os.Getenv("POSTGRES_PASSWORD"),
		PostgresName:     os.Getenv("POSTGRES_NAME"),
	}, nil
}
