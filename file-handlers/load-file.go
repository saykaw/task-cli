package filehandlers

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/saykaw/task-cli/types"
)

func LoadFile() []types.Task {
	file, err := os.OpenFile("tasks.json", os.O_RDWR|os.O_CREATE, 0766)
	if err != nil {
		fmt.Print(err.Error())
	}
	defer file.Close()

	loadedTasks := []types.Task{}

	data, err := os.ReadFile("tasks.json")
	if err != nil {
		fmt.Println("Error reading file:", err.Error())
	}

	if len(data) > 0 {
		err = json.Unmarshal(data, &loadedTasks)
		if err != nil {
			fmt.Println("Error unmarshalling json:", err.Error())
		}
	}
	return loadedTasks
}
