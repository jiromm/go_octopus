package octopus

import (
	"fmt"
	"github.com/jiromm/easyssh"
)

type Task struct {
	Name    string
	Command string
}

func (task *Task) Execute(ssh *easyssh.MakeConfig) (outStr string, err error) {
	fmt.Println("Executing '" + task.Command + "'")

	return ssh.Run(task.Command)
}
