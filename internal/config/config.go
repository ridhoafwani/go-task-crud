package config

import (
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

var AppConfig *Config

type Config struct {
	JWTSecretKey string
	Port         string
}

func LoadConfig() {
	secretKey := os.Getenv("JWT_SECRET_KEY")
	if secretKey == "" {
		log.Fatal("JWT_SECRET_KEY is not set in environment variables")
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT is not set in environment variables")
	}

	AppConfig = &Config{
		JWTSecretKey: secretKey,
		Port:         port,
	}
}
