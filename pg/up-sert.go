package pg

import (
	"database/sql"

	"github.com/Natchalit/gin-x-v1/sqlx"
)

func (c *Connect) UpSertBatch(
	table string,
	r *sqlx.Row,
	colsConflict []string, // คอลัมล์ที่อยู่ใน ON CONFLICT
	colsInsert []string, // คอลัมล์ที่จะทำการ insert
	colsUpdate []string, // คอลัมล์ที่อยู่ใน DO UPDATE SET
	colsExclude []string, // คอลัมล์ที่จะไม่ insert/update
	batchSize uint) (*[]sql.Result, error) {

	if ex := c._Connect(); ex != nil {
		return nil, ex
	}

	return c.db.UpSertBatch(table, r, colsConflict, colsInsert, colsUpdate, colsExclude, batchSize)
}
