package operationhandlers

import (
	"time"

	filehandlers "github.com/saykaw/task-cli/file-handlers"
)

func Update(id int, taskDescription string) {
	loadedTasks := filehandlers.LoadFile()
	for i, loadedTask := range loadedTasks {
		if loadedTask.Id == id {
			loadedTasks[i].TaskDescription = taskDescription
			loadedTasks[i].UpdatedAt = time.Now()
		}
	}
	filehandlers.SaveFile(loadedTasks)
}
