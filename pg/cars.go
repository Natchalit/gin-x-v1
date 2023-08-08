package pg

import (
	"github.com/Natchalit/gin-x-v1/connections"
	"github.com/Natchalit/gin-x-v1/sqlx"
)

func Cars() (*sqlx.Sqlx, error) {
	res, ex := connections.ConnectionSql(`dev_liyl`)
	if ex != nil {
		return nil, ex
	}
	return res, nil
}
