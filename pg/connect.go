package pg

import "github.com/Natchalit/gin-x-v1/connections"

func (c *Connect) _Connect() error {
	db, ex := connections.ConnectionSql(c.dbName)
	if ex != nil {
		return ex
	}

	c.db = db.Db
	return nil
}
