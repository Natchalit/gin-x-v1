package sqlx

import (
	"database/sql"
	"fmt"
	"strings"

	caseconvert "github.com/Natchalit/gin-x-v1/case-convert"
	"github.com/Natchalit/gin-x-v1/ginx"
)

/*
	INSERT INTO your_table (id, col1, col2, col3, col4, col5)
	VALUES
		(1, 'val1', 'val2', 'val3', 'val4', 'val5'),
		(2, 'val6', 'val7', 'val8', 'val9', 'val10')
	ON CONFLICT (id)
	DO UPDATE SET
		col1 = EXCLUDED.col1,
		col2 = EXCLUDED.col2,
		col3 = EXCLUDED.col3,
		col4 = EXCLUDED.col4,
		col5 = EXCLUDED.col5;
*/

func (db *DB) UpSertBatch(table string, r *Row, conflict []string, batchSize uint) (*[]sql.Result, error) {
	return ExecUpSert(db.Db, table, r, conflict, batchSize)
}

func ExecUpSert(db *sql.DB, table string, r *Row, conflict []string, batchSize uint) (*[]sql.Result, error) {

	rows := r.Rows

	if vRow := rows[0]; vRow == nil {
		return nil, ginx.BadRequest(`not found data`)
	}

	// recheck new columns
	r.Columns = []string{}
	for k := range rows[0] {
		r.Columns = append(r.Columns, caseconvert.ToSnake(k))
	}

	insertCol := fmt.Sprintf(`(%s)`, strings.Join(r.Columns, `,`))
	con_conflict := fmt.Sprintf(`(%s)`, strings.Join(conflict, `,`))

	excluded := ``
	for i, col := range r.Columns {
		excluded += fmt.Sprintf(`%s = EXCLUDED.%s`, col, col)
		if i+1 < len(r.Columns) {
			excluded += `,`
		}
	}

	val := ``
	resultx := []sql.Result{} // Slice to store results

	totalRows := len(rows)
	for i := 0; i < totalRows; i += int(batchSize) {
		end := i + int(batchSize)
		if end > totalRows {
			end = totalRows
		}

		batchRows := rows[i:end]

		val = "" // Reset val for each batch
		args := []interface{}{}

		for _, vRow := range batchRows {
			cols := len(vRow)
			buff := make([]string, cols)

			for i := 1; i <= cols; i++ {
				buff[i-1] = fmt.Sprintf(`$%d`, i)
				args = append(args, vRow[r.Columns[i-1]])
			}

			val += fmt.Sprintf(`(%s)`, strings.Join(buff, `,`))

			if i+1 < len(batchRows)-1 {
				val += ","
			}
		}

		query := fmt.Sprintf(`
            INSERT INTO %s %s
            VALUES %s
            ON CONFLICT %s
            DO UPDATE SET 
                %s`,
			table,
			insertCol,
			val,
			con_conflict,
			excluded)

		fmt.Println(query)
		result, ex := db.Exec(query, args...)
		if ex != nil {
			return nil, ex
		}

		resultx = append(resultx, result)
	}

	return &resultx, nil
}
