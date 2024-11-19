package utils

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"

	"github.com/lem3s/go-for-it/model"
)

func WriteToCsv(tasks []model.Task) {
	file, err := os.Create(TasksPath)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	writer := csv.NewWriter(file)

	defer writer.Flush()

	writer.Write([]string{"Id", "Description", "DateCreated", "IsDone"})

	writer.WriteAll(convertToStringMatrix(tasks))
}

func convertToStringMatrix(taskList []model.Task) [][]string {
	var stringMatrix [][]string

	for _, task := range taskList {
		var tempStringLine []string
		tempStringLine = append(tempStringLine, strconv.FormatInt(int64(task.Id), 10))
		tempStringLine = append(tempStringLine, task.Description)
		tempStringLine = append(tempStringLine, task.DateCreated.Format("2006-01-02 15:04:05"))
		tempStringLine = append(tempStringLine, strconv.FormatBool(task.IsDone))

		stringMatrix = append(stringMatrix, tempStringLine)
	}

	return stringMatrix
}
