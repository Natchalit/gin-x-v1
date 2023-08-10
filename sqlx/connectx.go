package sqlx

import (
	"database/sql"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/Natchalit/gin-x-v1/logx"
)

type DBX struct {
	dbName         string
	dbKey          string
	driver         string
	Db             *sql.DB
	_L_TABLE_EMPTY *sync.Mutex
}

var (
	_M_DBS   = map[string]*DBX{}
	_L_DBS   = &sync.Mutex{}
	POSTGRES = `postgres`
)

func Connectx(dbName, dsn string) (*DBX, error) {

	dbKey := fmt.Sprintf(`pg:%s`, dbName)
	db, ex := getInstance(os.Getpid(), dbKey, POSTGRES, dsn)
	if db != nil {
		db.dbName = dbName
	}

	return db, ex
}

func getInstance(pid int, dbKey, driver, dsn string) (*DBX, error) {

	// lock/unlock
	_L_DBS.Lock()
	defer _L_DBS.Unlock()

	dbKey = strings.ToUpper(dbKey)
	if _M_DBS[dbKey] != nil {
		return _M_DBS[dbKey], nil
	}

	db, ex := sql.Open(driver, dsn)
	if ex != nil {
		return nil, ex
	}

	logx.Warnf(`[%v:%v] db instance, Open !!!`, dbKey, pid)
	_M_DBS[dbKey] = &DBX{
		dbKey:          dbKey,
		driver:         driver,
		Db:             db,
		_L_TABLE_EMPTY: &sync.Mutex{},
	}

	return _M_DBS[dbKey], nil
}

func CleanConnections() {

	_L_DBS.Lock()
	defer _L_DBS.Unlock()
	for _, v := range _M_DBS {
		_ = v.Close()
	}

}
