package session

import (
	"fmt"

	t "../task"
	c "../config"
)

var I = 0

type Session struct {
	Name  string
	Tasks [10]*t.Task
	Config *c.Config
}

func (session *Session) AddTask(task *t.Task) {
	session.Tasks[I] = task
	I += 1

	fmt.Println("Task '" + task.Name + "' Has been added to session '" + session.Name + "'")
}

func (session *Session) Run() {
	for _, i := range session.Tasks {
		if i == nil {
			continue
		}

		i.Execute()
	}

	fmt.Println(session.Config)
}
