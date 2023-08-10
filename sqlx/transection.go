package sqlx

import (
	"context"

	"github.com/Natchalit/gin-x-v1/errorx"
)

func (db *DB) Transaction(callback func(*Tx) error) error {
	return db.TransactionContext(context.Background(), callback)
}

func (db *DB) TransactionContext(ctx context.Context, callback func(*Tx) error) error {
	// Start a new transaction
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		if r := recover(); r != nil {
			// rollback
			_ = tx.Rollback()
			// panic
			defer func() {
				panic(r)
			}()
		}
	}()

	// callback
	err = callback(tx)
	if err != nil {
		_ = tx.Rollback()
	} else {
		err = tx.Commit()
		if err != nil {
			err = errorx.Internal(`tx.Commit:%s`, err.Error())
		}
	}
	return err
}
