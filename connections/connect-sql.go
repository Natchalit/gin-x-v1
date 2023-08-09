package connections

import (
	"database/sql"
	"fmt"

	"github.com/Natchalit/gin-x-v1/ginx"
	"github.com/Natchalit/gin-x-v1/sqlx"
	_ "github.com/lib/pq"
)

func ConnectionSql(dbName string) (*sqlx.Sqlx, error) {

	// postgres://devnick:8FqDPl1daJCIyMffVWbzx9xC7sHl6dZt@dpg-cj85ijdjeehc73a6d9hg-a.singapore-postgres.render.com/dev_liyl

	PG_USER := `devnick`
	PG_PASS := `8FqDPl1daJCIyMffVWbzx9xC7sHl6dZt`
	PG_HOST := `dpg-cj85ijdjeehc73a6d9hg-a.singapore-postgres.render.com`
	PG_PORT := `5432`

	// Open a database connection

	dsn := fmt.Sprintf(
		`host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=UTC`,
		PG_HOST, PG_USER, PG_PASS, dbName, PG_PORT,
	)
	db, ex := sql.Open("postgres", dsn)
	if ex != nil {
		return nil, ginx.InternalServerError(`can not connect Postgres [%v]`, ex)
	}
	defer db.Close()

	return &sqlx.Sqlx{
		Db: db,
	}, nil
}
