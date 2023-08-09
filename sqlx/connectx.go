package sqlx

import (
	"database/sql"
	"sync"

	"github.com/Natchalit/gin-x-v1/ginx"
)

var (
	_M_DBS   = map[string]DB{}
	_L_DBS   = &sync.Mutex{}
	POSTGRES = `postgres`
)

func Connectx(dbName, dsn string) (*DB, error) {

	db, ex := sql.Open(POSTGRES, dsn)
	if ex != nil {
		return nil, ginx.InternalServerError(`can not connect Postgres [%v]`, ex)
	}

	return &DB{
		Db: db,
	}, nil
}

func CleanConnections() {

	_L_DBS.Lock()
	defer _L_DBS.Unlock()
	for _, v := range _M_DBS {
		_ = v.Close()
	}

}
