package db

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/lib/pq"
)

type DBParam struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

var (
	db         *sql.DB
	syncDBOnce sync.Once
)

func InitDB(param DBParam) (*sql.DB, error) {
	var err error

	syncDBOnce.Do(func() {
		if db == nil {
			dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
				param.User, param.Password, param.Host, param.Port, param.Name)
			db, err = sql.Open("postgres", dsn)

			err = db.Ping()

			if err != nil {
				return
			}
		}
	})
	return db, err
}
