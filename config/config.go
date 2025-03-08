package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using default vals")
	}
}

func GetEnv(key, fallback string) string {
	val, exists := os.LookupEnv(key)
	if !exists {
		return fallback
	}
	return val
}
