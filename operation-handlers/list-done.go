package operationhandlers

import (
	"fmt"

	filehandlers "github.com/saykaw/task-cli/file-handlers"
	"github.com/saykaw/task-cli/types"
)

func ListDone() {
	loadedTasks := filehandlers.LoadFile()
	doneTasks := []types.Task{}
	for _, loadedTask := range loadedTasks {
		if loadedTask.Status == "done" {
			doneTasks = append(doneTasks, loadedTask)
		}
	}
	fmt.Print("Your done tasks are:\n")
	fmt.Print(doneTasks)
}
