package main

import (
	"fmt"
	"os"
	"strconv"

	operationhandlers "github.com/saykaw/task-cli/operation-handlers"
)

func main() {

	fmt.Print(os.Args)
	switch os.Args[1] {
	case "add":
		operationhandlers.Add(os.Args[2])
	case "update":
		operationhandlers.Update(idArgToInt(os.Args[2]), os.Args[3])
	case "delete":
		operationhandlers.Delete(idArgToInt(os.Args[2]))
	case "list":
		if len(os.Args) > 2 {
			switch os.Args[2] {
			case "todo":
				operationhandlers.ListTodo()
			case "in-progress":
				operationhandlers.ListInProgress()
			case "done":
				operationhandlers.ListDone()
			}
		} else {
			operationhandlers.List()
		}
	case "mark-in-progress":
		operationhandlers.MarkInProgress(idArgToInt(os.Args[2]))
	case "mark-done":
		operationhandlers.MarkDone(idArgToInt(os.Args[2]))
	case "mark-todo":
		operationhandlers.MarkTodo(idArgToInt(os.Args[2]))
	}

}

func idArgToInt(arg string) int {
	id, err := strconv.Atoi(arg)
	if err != nil {
		fmt.Println("Invalid ID. Please provide a numeric ID.")
	}
	return id
}
