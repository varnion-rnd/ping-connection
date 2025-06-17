package pingconnection

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func PingMySQL(dsn string) error {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	defer db.Close()

	return db.Ping()
}
