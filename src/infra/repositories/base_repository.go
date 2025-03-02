package repositories

import (
	"app/src/infra/database"
)

var BaseRepository = database.NewRepository(
	"base",
	"id",
	[]string{"id", "name", "created_at", "updated_at"},
	[]string{"id", "name", "created_at"},
	[]string{"id", "name"},
	[]string{"name"},
)
