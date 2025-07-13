package env

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func loadEnv() bool {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return false
	}
	return true
}

func GetString(key string, fallback string) string {
	if !loadEnv() {
		return fallback
	}

	val, ok := os.LookupEnv(key)
	if ok {
		return val
	}
	return fallback
}

func GetInt(key string, fallback int) int {
	if !loadEnv() {
		return fallback
	}

	val, ok := os.LookupEnv(key)
	if ok {
		valAsInt, err := strconv.Atoi(val)
		if err != nil {
			return fallback
		}
		return valAsInt
	}
	return fallback
}
