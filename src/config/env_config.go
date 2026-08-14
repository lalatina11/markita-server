package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func check() error {
	return godotenv.Load()
}

func GetEnv(key string) string {
	err := check()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return os.Getenv(key)
}
