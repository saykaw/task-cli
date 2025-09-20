package filehandlers

import (
	"encoding/json"
	"os"

	"github.com/saykaw/task-cli/types"
)

func SaveFile(loadedTasks []types.Task) {
	content, err := json.Marshal(loadedTasks)
	if err != nil {
		println("Error marshalling json:", err.Error())
	}

	err = os.WriteFile("tasks.json", content, 0644)
	if err != nil {
		println("Error writing to file:", err.Error())
	}
}
