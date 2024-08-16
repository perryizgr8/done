package dbase

import (
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func Init() {
	db := getdb()
	defer db.Close()

	q := `
	CREATE TABLE IF NOT EXISTS tasks (id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, desc TEXT, done INTEGER);
	`
	_, err := db.Exec(q)
	if err != nil {
		log.Fatal(err)
	}
}
