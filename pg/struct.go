package pg

import (
	"github.com/Natchalit/gin-x-v1/sqlx"
)

type Connect struct {
	db     *sqlx.DB
	DBName string
}
