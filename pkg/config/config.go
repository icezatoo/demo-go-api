package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Environment string
	Port        string
	DbURL       string
	JWT_SECRET  string
}

const (
	postgresURL = "postgresql://%s:%s@%s/%s?sslmode=disable"
)

func LoadConfigENV() *Config {
	if err := godotenv.Load(); err != nil && os.Getenv("GO_ENV") != "production" {
		logrus.Error("Error loading .env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	port := os.Getenv("HOST_PORT")
	env := os.Getenv("GO_ENV")

	dbURL := fmt.Sprintf(postgresURL, dbUser, dbPass, dbHost, dbName)

	return &Config{
		Environment: env,
		Port:        port,
		DbURL:       dbURL,
		JWT_SECRET:  os.Getenv("JWT_SECRET"),
	}
}
