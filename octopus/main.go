package main

import (
	"fmt"
)

type Task struct {
	Name  string
	Power int
}

func main() {
	task := Task{"xxxx", 4}

	fmt.Println(task.Name)
}
