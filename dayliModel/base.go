package dayliModel

import (
	"database/sql"
)

var db *sql.DB

func InitDB(dataSourceName string) error {
	var err error

	db, err = sql.Open("sqlite3", dataSourceName)
	if err != nil {
		return err
	}
	return db.Ping()
}

func InitCreateTable() {
	CreateTableDayli()
}
