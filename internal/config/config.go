package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL  string
	DatabaseName string
	Port         string
}

func Load() (*Config, error) {
	var err error
	isProduction := os.Getenv("IS_PRODUCTION")

	if isProduction == "" {
		var err error = godotenv.Load()

		if err != nil {
			log.Println("Warning: .env file not found, using environment variables")
		}
	}

	var config *Config = &Config{
		DatabaseURL:  os.Getenv("MONGODB_URI"),
		DatabaseName: os.Getenv("DATABASE_NAME"),
		Port:         os.Getenv("PORT"),
	}

	return config, err
}
