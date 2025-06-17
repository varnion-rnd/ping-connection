package pingconnection

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func PingPostgres(dsn string) error {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return err
	}
	defer db.Close()

	return db.Ping()
}
