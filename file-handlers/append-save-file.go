package filehandlers

import (
	"encoding/json"
	"os"
	"time"

	"github.com/saykaw/task-cli/types"
)

func AppendAndSaveFile(taskDescription string, tasks []types.Task) {
	highestId := findHighestId(tasks)
	userInput := types.Task{Id: highestId + 1, TaskDescription: taskDescription, Status: "todo", CreatedAt: time.Now(), UpdatedAt: time.Now()}
	tasks = append(tasks, userInput)
	content, err := json.Marshal(tasks)
	if err != nil {
		println("Error marshalling json:", err.Error())
	}

	err = os.WriteFile("tasks.json", content, 0644)
	if err != nil {
		println("Error writing to file:", err.Error())
	}

}
