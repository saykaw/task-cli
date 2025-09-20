package operationhandlers

import (
	"fmt"

	filehandlers "github.com/saykaw/task-cli/file-handlers"
)

func List() {
	loadedTasks := filehandlers.LoadFile()
	fmt.Printf("Here are your tasks:\n")
	for _, task := range loadedTasks {
		fmt.Printf("ID: %d | Description: %s | Status: %s | Created Time: %v | Updated Time : %v |\n", task.Id, task.TaskDescription, task.Status, task.CreatedAt, task.UpdatedAt)
	}
}
