package database

import (
	"database/sql"
)

func ExecuteSQL(transaction *Transaction, sqlQuery string, args ...any) (sql.Result, error) {
	if transaction.transaction != nil {
		return transaction.transaction.Exec(sqlQuery, args...)
	}
	return transaction.databaseConnection.Exec(sqlQuery, args...)
}

func ExecuteQuery(transaction *Transaction, sqlQuery string, args ...any) (*sql.Rows, error) {
	if transaction.transaction != nil {
		return transaction.transaction.Query(sqlQuery, args...)
	}
	return transaction.databaseConnection.Query(sqlQuery, args...)
}
