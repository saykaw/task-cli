package operationhandlers

import (
	"fmt"

	filehandlers "github.com/saykaw/task-cli/file-handlers"
	"github.com/saykaw/task-cli/types"
)

func ListInProgress() {
	loadedTasks := filehandlers.LoadFile()
	inProgressTasks := []types.Task{}
	for _, loadedTask := range loadedTasks {
		if loadedTask.Status == "in-progress" {
			inProgressTasks = append(inProgressTasks, loadedTask)
		}
	}
	fmt.Print("Your in-progress tasks are:\n")
	fmt.Print(inProgressTasks)
}
