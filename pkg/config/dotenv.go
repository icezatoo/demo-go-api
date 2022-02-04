package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadConfig() {
	if err := godotenv.Load(); err != nil && os.Getenv("GO_ENV") != "production" {
		log.Fatal("Error loading .env file")
	}
}

func GetConfigByKey(key string) string {
	return os.Getenv(key)
}
