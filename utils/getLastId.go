package utils

func GetLastId() int {
	taskList := ReadFromCsv()

	if len(taskList) == 0 {
		return 0
	}

	return taskList[len(taskList)-1].Id
}
