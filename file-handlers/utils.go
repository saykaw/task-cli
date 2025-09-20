package filehandlers

import "github.com/saykaw/task-cli/types"

func findHighestId(tasks []types.Task) int {
	highestId := 0
	for _, task := range tasks {
		if task.Id > highestId {
			highestId = task.Id
		}
	}
	return highestId
}
