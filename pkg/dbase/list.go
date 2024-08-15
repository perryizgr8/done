package dbase

import (
	"log"

	"github.com/perryizgr8/done/pkg/common"

	_ "github.com/mattn/go-sqlite3"
)

func List() []common.Task {
	db := getdb()

	q := `
	SELECT id, desc, doneon FROM tasks ORDER BY id DESC;
	`
	rows, err := db.Query(q)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var tasks []common.Task
	for rows.Next() {
		var task common.Task
		err := rows.Scan(&task.Id, &task.Desc, &task.Doneon)
		if err != nil {
			log.Fatal(err)
		}
		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return tasks
}
