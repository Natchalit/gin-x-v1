package pg

import (
	"database/sql"

	"github.com/Natchalit/gin-x-v1/sqlx"
)

func (c *Connect) UpSertBatch(table string, r *sqlx.Row, conflict []string, batchSize uint) (*[]sql.Result, error) {
	return sqlx.ExecUpSert(c.db, table, r, conflict, batchSize)
}
