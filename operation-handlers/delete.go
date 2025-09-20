package operationhandlers

import (
	filehandlers "github.com/saykaw/task-cli/file-handlers"
	"github.com/saykaw/task-cli/types"
)

func Delete(id int) {
	loadedTasks := filehandlers.LoadFile()
	copyLoadedTasks := []types.Task{}
	for _, loadedTask := range loadedTasks {
		if loadedTask.Id != id {
			copyLoadedTasks = append(copyLoadedTasks, loadedTask)
		}
	}
	filehandlers.SaveFile(copyLoadedTasks)
}
