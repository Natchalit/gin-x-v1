package sqlx

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	caseconvert "github.com/Natchalit/gin-x-v1/case-convert"
	"github.com/Natchalit/gin-x-v1/errorx"
	"github.com/Natchalit/gin-x-v1/validx"
)

func (db *DB) Begin() (*Tx, error) {
	tx, ex := db.Db.Begin()
	if ex != nil {
		return nil, ex
	}

	return &Tx{
		Tx: tx,
	}, nil
}

func (db *DB) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	tx, ex := db.Db.BeginTx(ctx, opts)
	if ex != nil {
		return nil, ex
	}

	return &Tx{
		Tx: tx,
	}, nil
}

func (db *DB) Query(query string, args ...any) (*sql.Rows, error) {
	return db.Db.Query(query, args...)
}

func (db *DB) Close() error {
	return db.Db.Close()
}

func (db *DB) Exec(query string, args ...any) (sql.Result, error) {
	return db._ExecContext(context.Background(), query, args...)
}

func (db *DB) _ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
	for {
		resp, err := func() (sql.Result, error) {
			return db.Db.ExecContext(ctx, query, args...)
		}()
		return resp, err
	}
}

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

func (db *DB) _UpSert(
	table string,
	rows *Row,
	colsConflict []string, // คอลัมล์ที่อยู่ใน ON CONFLICT
	colsInsert []string, // คอลัมล์ที่จะทำการ insert
	colsUpdate []string, // คอลัมล์ที่อยู่ใน DO UPDATE SET
	colsExclude []string, // คอลัมล์ที่จะไม่ insert/update
	batchSize uint,
	exec func(string, *[]interface{}) (sql.Result, error),
) (*[]sql.Result, error) {

	// result
	results := []sql.Result{}

	// ไม่มีรายการที่จะ insert
	if len(rows.Rows) == 0 {
		return &results, nil
	}

	// validate
	if len(colsConflict) == 0 {
		return nil, errorx.Internal(`dxInsertUpdate, not found colsConflict`)
	}

	// คอลัมล์ที่จะทำการ insert
	if len(colsInsert) == 0 {
		// เอาทุกคอลัมล์
		colsInsert = rows.Columns
	} else {
		cols := []string{}
		for k := range colsInsert {
			if validx.IsContains(&rows.Columns, &colsInsert[k]) {
				cols = append(cols, colsInsert[k])
			}
			colsInsert = cols
		}
	}

	// ถ้าไม่ได้กำหนดคอลัมล์ udpate ให้เอาจาก rows.columns
	isDoUpdate := true
	if colsUpdate == nil {
		colsUpdate = append(colsUpdate, rows.Columns...)
	} else {
		if len(colsUpdate) == 0 {
			isDoUpdate = false // DO NOTHING
		}
	}

	// ตัดคอลัมล์ที่ไม่ต้องการออก
	if len(colsExclude) > 0 {
		// คอลัมล์ที่จะ insert
		colx := []string{}
		for i := range colsInsert {
			if validx.IsContains(&colsExclude, &colsInsert[i]) {
				continue // ยกเว้นคอลัมล์
			}
			colx = append(colx, colsInsert[i])
		}
		colsInsert = colx
		// คอลัมล์ที่จะ update
		if isDoUpdate {
			colx := []string{}
			for i := range colsUpdate {
				if validx.IsContains(&colsExclude, &colsUpdate[i]) {
					continue // ยกเว้นคอลัมล์
				}
				colx = append(colx, colsUpdate[i])
			}
			colsUpdate = colx
		}
	}

	// ตัดคอลัมล์มล์ conflict ออกจาก colsUpdate
	if isDoUpdate {
		colx := []string{}
		for i := range colsUpdate {
			if validx.IsContains(&colsConflict, &colsUpdate[i]) {
				continue // ยกเว้นคอลัมล์
			}
			colx = append(colx, colsUpdate[i])
		}
		colsUpdate = colx
	}

	// คอลัมล์ที่อยู่ใน colsUpdate ต้องอยู่ใน colsInsert ด้วย
	if isDoUpdate {
		colx := []string{}
		for i := range colsUpdate {
			if validx.IsContains(&colsInsert, &colsUpdate[i]) {
				colx = append(colx, colsUpdate[i])
			}
		}
		colsUpdate = colx
	}

	var parx, batchCount uint
	pars := []string{}
	vals := []interface{}{}

	// get command string
	getCommand := func() string {
		return fmt.Sprintf(`INSERT INTO %s (%s) VALUES (%s) ON CONFLICT (%s)`,
			table, strings.Join(colsInsert, `,`),
			strings.Join(pars, `),(`),
			strings.Join(colsConflict, `,`),
		)
	}

	// execute
	onExecute := func() error {
		cmd := getCommand()
		// ถ้าไม่่ได้กำหนดให้เอาทุกคอลัมล์
		if isDoUpdate && len(colsUpdate) > 0 {
			// DO UPDATE SET col1 = EXCLUDED.col1, col2 = EXCLUDED.col2, ...  (ทุกคอลัมล์ ยกเว้น conflict)
			colx := []string{}
			for _, v := range colsUpdate {
				colx = append(colx, fmt.Sprintf(`%s = EXCLUDED.%s`, v, v))
			}
			cmd += ` DO UPDATE SET ` + strings.Join(colx, `,`)
		} else {
			// DO NOTHING
			cmd += ` DO NOTHING`
		}
		result, ex := exec(cmd, &vals)
		if ex != nil {
			return ex
		}
		results = append(results, result)
		return nil
	}

	// loop
	for _, vr := range rows.Rows {
		parm := []string{}
		// คอลัมล์ที่จะ insert
		for _, vc := range colsInsert {
			parx++
			// if rows.DriverName == SQLSERVER {
			// 	parm = append(parm, fmt.Sprintf(`@p%v`, parx))
			// } else {
			parm = append(parm, fmt.Sprintf(`$%v`, parx))
			// }
			vals = append(vals, vr.Get(caseconvert.ToSnake(vc)))
		}
		pars = append(pars, strings.Join(parm, `,`))
		batchCount++
		if batchCount == batchSize {
			// execute
			if ex := onExecute(); ex != nil {
				return nil, ex
			}
			// clear
			batchCount = 0
			parx = 0
			pars = []string{}
			vals = []interface{}{}
		}
	}

	if batchCount > 0 {
		// execute
		if ex := onExecute(); ex != nil {
			return nil, ex
		}
	}

	return &results, nil
}
