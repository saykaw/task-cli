package operationhandlers

import (
	"fmt"

	filehandlers "github.com/saykaw/task-cli/file-handlers"
	"github.com/saykaw/task-cli/types"
)

func ListTodo() {
	loadedTasks := filehandlers.LoadFile()
	todoTasks := []types.Task{}
	for _, loadedTask := range loadedTasks {
		if loadedTask.Status == "todo" {
			todoTasks = append(todoTasks, loadedTask)
		}
	}
	fmt.Print("Your todo tasks are:\n")
	fmt.Print(todoTasks)
}
