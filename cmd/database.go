package cmd

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db *sql.DB
	// once sync.Once
)

func initDatabase(dbConf *Database) *sql.DB {
	dbConn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		dbConf.Username,
		dbConf.Password,
		dbConf.Host,
		dbConf.Port,
		dbConf.Name,
	)

	db, err := sql.Open(dbConf.Driver, dbConn)
	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}
