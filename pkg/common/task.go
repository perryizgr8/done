package common

import "time"

type Task struct {
	Id     int
	Desc   string
	Doneon time.Time
}
