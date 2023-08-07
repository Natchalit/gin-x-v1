package migratorx

import (
	"fmt"

	"github.com/Natchalit/gin-x-v1/connections"
	"github.com/Natchalit/gin-x-v1/logx"
)

func AutoMigrate(dbName string, dst ...any) error {
	db, ex := connections.ConnectionGorm(dbName)
	if ex != nil {
		return ex
	}

	if ex = db.AutoMigrate(dst...); ex != nil {
		return fmt.Errorf("AutoMigration [%v], %v", dbName, ex)
	} else {
		logx.Infof("AutoMigration success. [%v]", dbName)
	}

	return nil
}
