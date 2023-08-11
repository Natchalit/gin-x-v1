package pg

import (
	"database/sql"
)

func (c *Connect) Exec(query string, args ...any) (sql.Result, error) {
	if ex := c._Connect(); ex != nil {
		return nil, ex
	}

	return c.db.Exec(query, args...)
}
