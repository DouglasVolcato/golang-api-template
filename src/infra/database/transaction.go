package database

import (
	"database/sql"
	"fmt"
)

type Transaction struct {
	transaction        *sql.Tx
	databaseConnection *sql.DB
}

func (transaction *Transaction) BeginTransaction() error {
	var err error
	transaction.transaction, err = transaction.databaseConnection.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	return nil
}

func (transaction *Transaction) CommitTransaction() error {
	if transaction.transaction == nil {
		return fmt.Errorf("transaction has not been started")
	}
	err := transaction.transaction.Commit()
	if err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}
	return nil
}

func (transaction *Transaction) RollbackTransaction() error {
	if transaction.transaction == nil {
		return fmt.Errorf("transaction has not been started")
	}
	err := transaction.transaction.Rollback()
	if err != nil {
		return fmt.Errorf("failed to rollback transaction: %w", err)
	}
	return nil
}
