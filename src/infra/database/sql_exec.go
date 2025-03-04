package database

import (
	"database/sql"
)

func ExecuteSQL(transaction *Transaction, sqlQuery string, args ...any) error {
	if transaction.transaction != nil {
		_, err := transaction.transaction.Exec(sqlQuery, args...)
		return err
	} else {
		_, err := transaction.databaseConnection.Exec(sqlQuery, args...)
		return err
	}
}

func ExecuteQuery(transaction *Transaction, sqlQuery string, args ...any) (*sql.Rows, error) {
	if transaction.transaction != nil {
		return transaction.transaction.Query(sqlQuery, args...)
	}
	return transaction.databaseConnection.Query(sqlQuery, args...)
}
