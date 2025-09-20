package operationhandlers

import filehandlers "github.com/saykaw/task-cli/file-handlers"

func Add(taskDescription string) {

	loadedTasks := filehandlers.LoadFile()
	filehandlers.AppendAndSaveFile(taskDescription, loadedTasks)

}
