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
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Print("Error converting id to int:", err.Error())
		}
		operationhandlers.Update(id, os.Args[3])
	case "delete":
		fmt.Print("delete")
	case "list":
		operationhandlers.List()

	}

}
