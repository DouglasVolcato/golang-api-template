package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

var (
	globalDatabaseConnection *sql.DB
	once                     sync.Once
)

func InitializeDatabaseConnection() *sql.DB {
	once.Do(func() {
		err := godotenv.Load()
		if err != nil {
			log.Println("No .env file found, using environment variables.")
		}

		databaseConnectionString := fmt.Sprintf(
			"postgres://%s:%s@%s:%s/%s?sslmode=disable",
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_NAME"),
		)

		globalDatabaseConnection, err = sql.Open("pgx", databaseConnectionString)
		if err != nil {
			log.Fatalf("Error connecting to the database: %v", err)
		}

		if err := globalDatabaseConnection.Ping(); err != nil {
			log.Fatalf("Database is not reachable: %v", err)
		}

		fmt.Println("Successfully connected to the database")
	})

	return globalDatabaseConnection
}
