package repositories

import (
	"app/src/infra/database"
	"fmt"

	"database/sql"
)

var BaseRepository = baseRepository()

type repositoryType struct {
	*database.Repository
}

func baseRepository() *repositoryType {
	return &repositoryType{
		Repository: database.NewRepository(
			"base_table",
			"id",
			[]string{"id", "name", "created_at", "updated_at"},
			[]string{"id", "name", "created_at"},
			[]string{"id", "name"},
			[]string{"name"},
		),
	}
}

func (repo *repositoryType) SelectCountAll(transaction *database.Transaction, limit int, offset int) (*sql.Rows, error) {
	return database.ExecuteQuery(
		transaction,
		fmt.Sprintf(
			"SELECT *, (SELECT COUNT(*) FROM base_table) FROM base_table LIMIT %d OFFSET %d",
			limit,
			offset,
		),
	)
}
