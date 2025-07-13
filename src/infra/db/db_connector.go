package db

import (
	"api-template/src/utils"
	"database/sql"
	"log"
	"strconv"

	_ "github.com/lib/pq"
)

var GlobalDbConnection *sql.DB

func Connect() {
	env := utils.Env{}
	host := env.GetString("DB_HOST", "localhost")
	port := env.GetInt("DB_PORT", 5432)
	user := env.GetString("DB_USER", "postgres")
	password := env.GetString("DB_PASSWORD", "password")
	dbname := env.GetString("DB_NAME", "api_db")
	connectionString := "host=" + host + " port=" + strconv.Itoa(port) + " user=" + user + " password=" + password + " dbname=" + dbname + " sslmode=disable"
	GlobalDbConnection, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	if err = GlobalDbConnection.Ping(); err != nil {
		log.Fatal(err)
	}
}

func Close() {
	err := GlobalDbConnection.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func BeginTransaction() *sql.Tx {
	tx, err := GlobalDbConnection.Begin()
	if err != nil {
		log.Fatal(err)
	}
	return tx
}

func CommitTransaction(tx *sql.Tx) {
	err := tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
}

func RollbackTransaction(tx *sql.Tx) {
	err := tx.Rollback()
	if err != nil {
		log.Fatal(err)
	}
}
