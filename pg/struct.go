package pg

import (
	"database/sql"
)

type Connect struct {
	db     *sql.DB
	DBName string
}
