package dbase

import (
	"log"

	"github.com/perryizgr8/done/pkg/common"

	_ "github.com/mattn/go-sqlite3"
)

func Add(task common.Task) error {
	log.Println("adding:", task.Desc)
	db := getdb()
	defer db.Close()

	txn, err := db.Begin()
	if err != nil {
		log.Fatal(err)
		return err
	}
	stmt, err := txn.Prepare("INSERT INTO tasks (desc, done) VALUES (?, ?)")
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(task.Desc, task.Done.UnixNano())
	if err != nil {
		log.Fatal(err)
		return err
	}
	err = txn.Commit()
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
