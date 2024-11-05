package cmd

import "time"

type Task struct {
	id          int
	description string
	dateCreated time.Time
	status      bool
}
