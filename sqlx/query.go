package sqlx

import (
	"database/sql"
	"strings"

	"github.com/Natchalit/gin-x-v1/logx"
)

func (db *DB) QueryScan(query string, args ...any) (*Row, error) {

	defer db.Db.Close()

	rows, ex := db.Db.Query(query, args...)
	if ex != nil {
		return nil, ex
	}

	return ScanData(rows)
}

func ScanData(rows *sql.Rows) (*Row, error) {

	defer rows.Close()

	// get cols name
	cols, _ := rows.Columns()
	rowx := []Map{}
	colsTypex := []ColumnType{}
	chkCols := map[string]bool{}
	// map data
	for rows.Next() {

		values := make([]interface{}, len(cols))
		row := make(map[string]interface{})

		for i := range cols {
			values[i] = new(interface{})
		}

		colsTypes, _ := rows.ColumnTypes()
		colsTypem := map[string]string{}

		for _, v := range colsTypes {
			if ok := chkCols[v.Name()]; !ok {

				colTypex := ColumnType{
					Name:             v.Name(),
					DatabaseTypeName: strings.ToUpper(v.DatabaseTypeName()),
				}
				colsTypex = append(colsTypex, colTypex)
				colsTypem[colTypex.Name] = colTypex.DatabaseTypeName
				chkCols[v.Name()] = true
			}
		}
		// scan
		if ex := rows.Scan(values...); ex != nil {
			logx.Errorf(ex.Error())
		}

		for i, column := range cols {
			val := *(values[i].(*interface{}))
			row[column] = val
		}

		rowx = append(rowx, row)
	}

	res := Row{
		Rows:        rowx,
		Columns:     cols,
		ColumnTypes: colsTypex,
	}

	return &res, nil
}
