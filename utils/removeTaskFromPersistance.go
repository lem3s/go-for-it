package utils

import (
	"github.com/lem3s/go-for-it/model"
)

func RemoveTask(task model.Task) {
	taskList := ReadFromCsv()

	taskList = append(taskList, task)

	WriteToCsv(taskList)
}
