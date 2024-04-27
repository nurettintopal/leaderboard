package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config func to get env value
func Config(key string) string {
	// Loads values from .env into the system
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env file")
	}
	return os.Getenv(key)
}
