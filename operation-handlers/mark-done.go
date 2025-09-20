package operationhandlers

import (
	"time"

	filehandlers "github.com/saykaw/task-cli/file-handlers"
)

func MarkDone(id int) {
	loadedTasks := filehandlers.LoadFile()
	for i, loadedTask := range loadedTasks {
		if loadedTask.Id == id {
			loadedTasks[i].Status = "done"
			loadedTasks[i].UpdatedAt = time.Now()
		}
	}
	filehandlers.SaveFile(loadedTasks)
}
