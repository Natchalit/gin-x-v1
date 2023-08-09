package sqlx

func (db *DB) Close() error {
	return db.Db.Close()
}
