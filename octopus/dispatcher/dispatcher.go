package dispatcher

import (
	"fmt"

	t "../task"
	c "../config"
)

var I = 0

type Dispatcher struct {
	Name  string
	Tasks [10]*t.Task
	Config *c.Config
}

func (dispatcher *Dispatcher) AddTask(task *t.Task) {
	dispatcher.Tasks[I] = task
	I += 1

	fmt.Println("Task '" + task.Name + "' has been added to dispatcher '" + dispatcher.Name + "'")
}

func (dispatcher *Dispatcher) Run() {
	for _, i := range dispatcher.Tasks {
		if i == nil {
			continue
		}

		i.Execute()
	}

	fmt.Println(dispatcher.Config)
}
