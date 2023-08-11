package sqlx

import (
	"database/sql"

	caseconvert "github.com/Natchalit/gin-x-v1/case-convert"
)

func (db *DB) UpSertBatch(table string,
	r *Row,
	colsConflict,
	colsInsert,
	colsUpdate,
	colsExclude []string,
	batchSize uint) (*[]sql.Result, error) {
	return db.ExecUpSert(table, r, colsConflict, colsInsert, colsUpdate, colsExclude, batchSize)
}

func (db *DB) ExecUpSert(
	table string,
	r *Row,
	colsConflict []string, // คอลัมล์ที่อยู่ใน ON CONFLICT
	colsInsert []string, // คอลัมล์ที่จะทำการ insert
	colsUpdate []string, // คอลัมล์ที่อยู่ใน DO UPDATE SET
	colsExclude []string, // คอลัมล์ที่จะไม่ insert/update
	batchSize uint,
) (*[]sql.Result, error) {

	rows := r.Rows
	lenCol := len(r.Columns)

	if (lenCol == 0) || (r.LenCols() != lenCol) {
		r.Columns = []string{}
		// recheck new columns
		for k := range rows[0] {
			r.Columns = append(r.Columns, caseconvert.ToSnake(k))
		}
	}

	resultx := []sql.Result{} // Slice to store results

	totalRows := len(rows)

	// for i := 0; i < totalRows; i += int(batchSize) {
	if float64(totalRows)/float64(batchSize) > 1 {
		if ex := db.Transaction(func(tx *Tx) error {
			_, ex := db._UpSert(table, r, colsConflict, colsInsert, colsUpdate, colsExclude, batchSize, func(query string, args *[]interface{}) (sql.Result, error) {
				return tx.Exec(query, *args...)
			})
			if ex != nil {
				return ex
			}
			return nil
		}); ex != nil {
			return nil, ex
		}
	} else {
		return db._UpSert(table, r, colsConflict, colsInsert, colsUpdate, colsExclude, batchSize, func(query string, args *[]interface{}) (sql.Result, error) {
			return db.Exec(query, *args...)
		})

	}

	return &resultx, nil
}
