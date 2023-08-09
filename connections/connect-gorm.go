package connections

import (
	"fmt"

	getenv "github.com/Natchalit/gin-x-v1/get-env"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectionGorm(dbName string) (*gorm.DB, error) {

	PG_HOST := getenv.Get(`PG_HOST`)
	// PG_PORT := getenv.Get(`PG_PORT`)
	PG_USER := getenv.Get(`PG_USER`)
	PG_PASS := getenv.Get(`PG_PASS`)

	dsn := fmt.Sprintf(
		`postgres://%s:%s@%s/%s`,
		PG_USER, PG_PASS, PG_HOST, dbName,
	)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
