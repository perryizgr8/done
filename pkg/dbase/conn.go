package dbase

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func getdb() *sql.DB {
	db, err := sql.Open("sqlite3", "./tasksdb/tasks.sqlite3")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	return db
}
