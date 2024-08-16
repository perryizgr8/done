package common

import "time"

type Task struct {
	Id   int
	Desc string
	Done time.Time
}

func (t Task) ShortDoneDate() string {
	return t.Done.Format(time.RFC822)
}
