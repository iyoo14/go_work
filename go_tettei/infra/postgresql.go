package infra

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func Connect(dsn string) *sql.DB {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return db
}
