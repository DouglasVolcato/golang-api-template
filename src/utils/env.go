package utils

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Env struct {
}

func (e *Env) loadEnv() bool {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return false
	}
	return true
}

func (e *Env) GetString(key string, fallback string) string {
	if !e.loadEnv() {
		return fallback
	}

	val, ok := os.LookupEnv(key)
	if ok {
		return val
	}
	return fallback
}

func (e *Env) GetInt(key string, fallback int) int {
	if !e.loadEnv() {
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
