package pg

import "github.com/Natchalit/gin-x-v1/sqlx"

func (c *Connect) QueryScan(query string, args ...any) (*sqlx.Row, error) {
	if ex := c._Connect(); ex != nil {
		return nil, ex
	}

	rows, ex := c.db.Query(query, args...)
	if ex != nil {
		return nil, ex
	}

	return sqlx.ScanData(rows)
}
