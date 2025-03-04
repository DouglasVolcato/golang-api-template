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

func NewRepository(tableName string, idField string, fields []string, publicFields []string, insertFields []string, updateFields []string) *Repository {
	return &Repository{tableName: tableName, idField: idField, fields: fields, publicFields: publicFields, insertFields: insertFields, updateFields: updateFields}
}

func (repo *Repository) Insert(transaction *Transaction, values []any) error {
	placeholders := make([]string, len(repo.insertFields))
	for i := range placeholders {
		placeholders[i] = fmt.Sprintf("$%d", i+1)
	}

	sqlQuery := fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES (%s)",
		repo.tableName,
		strings.Join(repo.insertFields, ", "),
		strings.Join(placeholders, ", "),
	)

	return ExecuteSQL(transaction, sqlQuery, values...)
}

func (repo *Repository) Update(transaction *Transaction, id string, values []any) error {
	placeholders := make([]string, len(repo.updateFields))
	for i := range placeholders {
		placeholders[i] = fmt.Sprintf("%s = $%d", repo.updateFields[i], i+1)
	}

	sqlQuery := fmt.Sprintf(
		"UPDATE %s SET %s WHERE %s = $%d",
		repo.tableName,
		strings.Join(placeholders, ", "),
		repo.idField,
		len(repo.updateFields)+1,
	)

	return ExecuteSQL(transaction, sqlQuery, append(values, id)...)
}

func (repo *Repository) Delete(transaction *Transaction, id string) error {
	return ExecuteSQL(transaction, fmt.Sprintf("DELETE FROM %s WHERE %s = ?", repo.tableName, repo.idField), id)
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
