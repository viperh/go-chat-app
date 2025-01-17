package config

import (
	"os"
)

type Config struct {
	DbUser     string
	DbPassword string
	DbHost     string
	DbPort     string
	DbName     string
	DbSslMode  string

	ServerPort string
	LogLevel   string
	JwtKey     string
}

func loadFromEnv() *Config {
	return &Config{
		DbUser:     os.Getenv("DB_USER"),
		DbPassword: os.Getenv("DB_PASSWORD"),
		DbHost:     os.Getenv("DB_HOST"),
		DbPort:     os.Getenv("DB_PORT"),
		DbName:     os.Getenv("DB_NAME"),
		DbSslMode:  os.Getenv("DB_SSL_MODE"),
		LogLevel:   os.Getenv("LOG_LEVEL"),
		ServerPort: os.Getenv("SERVER_PORT"),
		JwtKey:     os.Getenv("JWT_KEY"),
	}
}

func NewConfig(mode string) *Config {
	return loadFromEnv()
}
