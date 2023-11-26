package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB_PORT string
	DB_NAME string
	DB_USER string
	DB_HOST string
	DB_PASS string
}

func NewConfig() (*Config, error) {
	err := godotenv.Load("../.env")
	if err != nil {
		return nil, err
	}

	return &Config{
		DB_PORT: os.Getenv("DB_PORT"),
		DB_NAME: os.Getenv("DB_NAME"),
		DB_USER: os.Getenv("DB_USER"),
		DB_HOST: os.Getenv("DB_HOST"),
		DB_PASS: os.Getenv("DB_PASS"),
	}, nil
}
