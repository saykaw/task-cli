package operationhandlers

import (
	"fmt"
	"strconv"
)

func IdArgToInt(arg string) int {
	id, err := strconv.Atoi(arg)
	if err != nil {
		fmt.Println("Invalid ID. Please provide a numeric ID.")
	}
	return id
}
