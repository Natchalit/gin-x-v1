package pg

import (
	"github.com/Natchalit/gin-x-v1/connections"
	"github.com/Natchalit/gin-x-v1/sqlx"
)

func (c *Connect) _Connect() error {
	db, ex := connections.ConPg(c.DBName)
	if ex != nil {
		return ex
	}
	c.db = &sqlx.DB{
		Db: db.Db,
	}
	return nil
}
