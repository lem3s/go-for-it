package model

import "time"

type Task struct {
	Id          int
	Description string
	DateCreated time.Time
	IsDone      bool
}
