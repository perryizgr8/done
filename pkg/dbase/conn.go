package dbase

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func getdb() *sql.DB {
	db, err := sql.Open("sqlite3", "./tasksdb/tasks.sqlite3")
	if err != nil {
		log.Printf("getdb: %v", err)
		panic(err)
	}
	return db
}
