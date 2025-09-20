package operationhandlers

import (
	"time"

	filehandlers "github.com/saykaw/task-cli/file-handlers"
)

func MarkTodo(id int) {
	loadedTasks := filehandlers.LoadFile()
	for i, loadedTask := range loadedTasks {
		if loadedTask.Id == id {
			loadedTasks[i].Status = "todo"
			loadedTasks[i].UpdatedAt = time.Now()
		}
	}
	filehandlers.SaveFile(loadedTasks)
}
