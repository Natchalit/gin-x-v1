package sqlx

import (
	"context"
	"database/sql"
)

func (tx *Tx) Rollback() error {
	return tx.Tx.Rollback()
}

func (tx *Tx) Commit() error {
	return tx.Tx.Commit()
}

func (tx *Tx) Exec(query string, args ...any) (sql.Result, error) {
	return tx._ExecContext(context.Background(), query, args...)
}

func (tx *Tx) _ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
	return tx.Tx.ExecContext(ctx, query, args...)
}
