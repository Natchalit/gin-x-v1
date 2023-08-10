package connections

import (
	"fmt"

	_ "github.com/lib/pq"

	getenv "github.com/Natchalit/gin-x-v1/get-env"
	"github.com/Natchalit/gin-x-v1/sqlx"
)

var ()

func ConnectionSql(dbName string) (*sqlx.DBX, error) {
	// postgres://devnick:8FqDPl1daJCIyMffVWbzx9xC7sHl6dZt@dpg-cj85ijdjeehc73a6d9hg-a.singapore-postgres.render.com/dev_liyl

	PG_USER := getenv.Get(`PG_USER`)
	PG_PASS := getenv.Get(`PG_PASS`)
	PG_HOST := getenv.Get(`PG_HOST`)
	PG_PORT := getenv.Get(`PG_PORT`)

	// Open a database connection

	dsn := fmt.Sprintf(
		`host=%v user=%v password=%v dbname=%v port=%v TimeZone=UTC`,
		PG_HOST, PG_USER, PG_PASS, dbName, PG_PORT,
	)
	return sqlx.Connectx(dbName, dsn)
}
