package sqlx

import (
	"database/sql"
	"fmt"
	"strings"
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

func (db *Sqlx) UpSert(table string, r *Row, conflict []string, Callback ValidationCallback) (*[]sql.Result, error) {

	defer db.Db.Close()

	if Callback != nil {
		for _, row := range r.Rows {
			if err := Callback(row); err != nil {
				return nil, err
			}
		}
	}

	insertCol := fmt.Sprintf(`(%s)`, strings.Join(r.Columns, `,`))
	con_confilct := fmt.Sprintf(`(%s)`, strings.Join(conflict, `,`))

	excluded := ``
	for i, col := range r.Columns {
		excluded += fmt.Sprintf(`%s = EXCLUDED.%s`, col, col)
		if i+1 < len(r.Columns) {
			excluded += `,`
		}
	}

	val := ``
	args := []interface{}{}
	for i, vRow := range r.Rows {
		cols := len(vRow)
		buff := make([]string, cols)

		for i := 1; i <= cols; i++ {
			buff[i-1] = fmt.Sprintf(`$%d`, i)
			args = append(args, vRow[r.Columns[i-1]])
		}

		val += fmt.Sprintf(`(%s)`, strings.Join(buff, `,`))

		i += 1
		if i < len(r.Rows) {
			val += `,`
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
		con_confilct,
		excluded)

	fmt.Println(query)
	result, ex := db.Db.Exec(query, args...)
	if ex != nil {
		return nil, ex
	}

	resultx := []sql.Result{}
	resultx = append(resultx, result)

	return &resultx, nil
}
