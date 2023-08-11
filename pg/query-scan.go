package pg

import (
	"context"

	"github.com/Natchalit/gin-x-v1/sqlx"
)

func (c *Connect) QueryScan(query string, args ...any) (*sqlx.Row, error) {
	return c.QueryContext(query, args...)
}

func (c *Connect) QueryContext(query string, args ...any) (*sqlx.Row, error) {

	if ex := c._Connect(); ex != nil {
		return nil, ex
	}

	rows, ex := c.db.Db.QueryContext(context.Background(), query, args...)
	if ex != nil {
		return nil, ex
	}

	return sqlx.ScanData(rows)
}
