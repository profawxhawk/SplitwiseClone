package configs

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func getEnvWithKey(key string) string {
	return os.Getenv(key)
}

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading env file")
	}
}
