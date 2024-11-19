package utils

import (
	"encoding/csv"
	"errors"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/lem3s/go-for-it/model"
)

const TasksPath = "data/tasks.csv"

func ReadFromCsv() []model.Task {
	if _, err := os.Stat(TasksPath); errors.Is(err, os.ErrNotExist) {
		var emptyTaskList []model.Task
		return emptyTaskList
	}

	file, err := os.Open(TasksPath)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	csvReader := csv.NewReader(file)

	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	return convertToTaskList(data)
}

func convertToTaskList(data [][]string) []model.Task {
	var tasks []model.Task
	for _, row := range data[1:] {
		var task model.Task
		task.Id, _ = strconv.Atoi(row[0])
		task.Description = row[1]
		task.DateCreated, _ = time.Parse("2006-01-02 15:04:05", row[2])
		task.IsDone, _ = strconv.ParseBool(row[3])
		tasks = append(tasks, task)
	}
	return tasks
}
