package database

import (
	"database/sql"
	"fmt"
	"strings"
)

type Repository struct {
	tableName    string
	idField      string
	fields       []string
	publicFields []string
	insertFields []string
	updateFields []string
}

func (repo *Repository) Insert(transaction *Transaction, values []any) {
	ExecuteSQL(transaction, fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", repo.tableName, strings.Join(repo.insertFields, ", "), strings.Repeat("?", len(repo.insertFields))), values...)
}

func (repo *Repository) Update(transaction *Transaction, values []any) {
	ExecuteSQL(transaction, fmt.Sprintf("UPDATE %s SET %s WHERE %s = ?", repo.tableName, strings.Join(repo.updateFields, ", "), repo.idField), values...)
}

func (repo *Repository) Delete(transaction *Transaction, id string) {
	ExecuteSQL(transaction, fmt.Sprintf("DELETE FROM %s WHERE %s = ?", repo.tableName, repo.idField), id)
}

func (repo *Repository) Select(transaction *Transaction, id string) (*sql.Rows, error) {
	return ExecuteQuery(transaction, fmt.Sprintf("SELECT %s FROM %s WHERE %s = ?", strings.Join(repo.fields, ", "), repo.tableName, repo.idField), id)
}

func (repo *Repository) SelectAll(transaction *Transaction, limit int, offset int) (*sql.Rows, error) {
	return ExecuteQuery(transaction, fmt.Sprintf("SELECT %s FROM %s LIMIT %d OFFSET %d", strings.Join(repo.publicFields, ", "), repo.tableName, limit, offset))
}

func (repo *Repository) ExecuteQuery(transaction *Transaction, sqlQuery string, args ...any) (*sql.Rows, error) {
	return ExecuteQuery(transaction, sqlQuery, args...)
}
