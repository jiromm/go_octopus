package main

import (
	"fmt"
)

type Task struct {
	Name    string
	Command string
}

func main() {
	task := &Task{
		Name:    "xxxx",
		Command: "cd /var/www && ls -la",
	}

	Super(task)

	fmt.Println(task)
}

func Super(s *Task) {
	s.Command += " | grep php"
}
