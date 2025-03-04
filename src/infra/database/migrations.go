package database

import (
	"app/src/domain/utils"
	"database/sql"
	"os"
)

func ExecuteDatabaseMigrations(globalDatabaseConnection *sql.DB) error {
	_, err := globalDatabaseConnection.Exec(`
		CREATE TABLE IF NOT EXISTS migrations (
			id VARCHAR(255) PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			created_at TIMESTAMP NOT NULL DEFAULT NOW()
		);
	`)
	if err != nil {
		return err
	}

	files, err := os.ReadDir("db/migrations")
	if err != nil {
		return err
	}

	var existingMigrations []string
	rows, err := globalDatabaseConnection.Query("SELECT name FROM migrations")
	if err != nil {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return err
		}
		existingMigrations = append(existingMigrations, name)
	}

	for _, file := range files {
		if !contains(existingMigrations, file.Name()) {
			_, err = globalDatabaseConnection.Exec(readFile("db/migrations/" + file.Name()))
			if err != nil {
				return err
			}
			_, err = globalDatabaseConnection.Exec(
				"INSERT INTO migrations (id, name) VALUES ($1, $2)",
				utils.GenerateUuid(),
				file.Name(),
			)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func readFile(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(data)
}
