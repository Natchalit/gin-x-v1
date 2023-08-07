package pg

import (
	"github.com/Natchalit/gin-x-v1/connections"
	"github.com/Natchalit/gin-x-v1/sqlx"
)

func Cars() (*sqlx.Sqlx, error) {
	return connections.ConnectionSql(`dev_liyl`)
}
