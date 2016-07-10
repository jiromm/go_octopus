package task

import (
	"fmt"
)

type Task struct {
	Name    string
	Command string
}

func (task *Task) Execute() {
	fmt.Println("Executing '" + task.Command + "'")
}
