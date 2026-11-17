package utils

import (
	"log"

	env "github.com/joho/godotenv"
)

func LoadEnv() {
	if err := env.Load(); err != nil {
		log.Fatal(".env file not found")
	}
}
