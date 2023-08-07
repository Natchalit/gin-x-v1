package connections

import (
	"database/sql"
	"fmt"

	"github.com/Natchalit/gin-x-v1/sqlx"
)

func ConnectionSql(dbName string) (*sqlx.Sqlx, error) {

	// postgres://devnick:8FqDPl1daJCIyMffVWbzx9xC7sHl6dZt@dpg-cj85ijdjeehc73a6d9hg-a.singapore-postgres.render.com/dev_liyl

	PG_USER := `devnick`
	PG_PASS := `8FqDPl1daJCIyMffVWbzx9xC7sHl6dZt`
	PG_HOST := `dpg-cj85ijdjeehc73a6d9hg-a.singapore-postgres.render.com`
	// PG_PORT := `5432`
	// DATABASE := `dev_liyl`

	connStr := fmt.Sprintf("postgres://%s:%s@%s/%s", PG_USER, PG_PASS, PG_HOST, dbName)
	db, ex := sql.Open("postgres", connStr)
	if ex != nil {
		return nil, ex
	}
	defer db.Close()

	return &sqlx.Sqlx{
		Db: db,
	}, nil
}
