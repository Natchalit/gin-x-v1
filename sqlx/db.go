package sqlx

func (db *DBX) Close() error {
	return db.Db.Close()
}
