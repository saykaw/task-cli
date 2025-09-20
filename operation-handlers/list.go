package operationhandlers

import (
	"fmt"

	filehandlers "github.com/saykaw/task-cli/file-handlers"
)

func List() {
	loadedTasks := filehandlers.LoadFile()
	fmt.Printf("Here are your tasks:\n")
	fmt.Println(loadedTasks)
}
