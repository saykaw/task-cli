package main

import (
	"fmt"
	"os"

	operationhandlers "github.com/saykaw/task-cli/operation-handlers"
)

func main() {

	fmt.Print(os.Args)
	switch os.Args[1] {
	case "add":
		operationhandlers.Add(os.Args[2])
	case "update":
		fmt.Printf("update")
	case "delete":
		fmt.Print("delete")
	case "list":
		operationhandlers.List()

	}

}
