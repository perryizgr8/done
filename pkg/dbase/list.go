package dbase

import (
	"log"
	"time"

	"github.com/perryizgr8/done/pkg/common"

	_ "github.com/mattn/go-sqlite3"
)

func List() []common.Task {
	db := getdb()
	defer db.Close()

	q := `
	SELECT id, desc, done FROM tasks ORDER BY id DESC;
	`
	rows, err := db.Query(q)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var tasks []common.Task
	for rows.Next() {
		var task common.Task
		var doneon int64
		err := rows.Scan(&task.Id, &task.Desc, &doneon)
		log.Printf("doneone scanned: %d", doneon)
		if err != nil {
			log.Println(err)
			continue
		}
		task.Done = time.Unix(0, doneon)
		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return tasks
}
