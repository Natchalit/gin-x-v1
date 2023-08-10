package sqlx

import (
	"database/sql"
)

type Map map[string]any

type DB struct {
	Db *sql.DB
}
type ColumnType struct {
	Name             string `json:"name"`
	DatabaseTypeName string `json:"database_type_name"`
}

type Row struct {
	Rows        []Map        `json:"rows"`
	Columns     []string     `json:"columns"`
	ColumnTypes []ColumnType `json:"columns_type"`
}

type ValidationCallback func(row Map) error
