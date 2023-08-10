package sqlx

import (
	"context"
	"database/sql"
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

func (db *DB) Close() error {
	return db.Db.Close()
}
