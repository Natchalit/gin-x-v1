package pg

import (
	"database/sql"
)

type Connect struct {
	db     *sql.DB
	dbName string
}
