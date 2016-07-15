package octopus

import (
	"fmt"
)

var I = 0

type Dispatcher struct {
	Name  string
	Tasks [10]*Task
	Config *Config
}

func (dispatcher *Dispatcher) AddTask(task *Task) {
	dispatcher.Tasks[I] = task
	I += 1

	fmt.Println("Task '" + task.Name + "' has been added to dispatcher '" + dispatcher.Name + "'")
}

func (dispatcher *Dispatcher) Run() {
	connector := Connector{
		Config: dispatcher.Config,
	}

	c := connector.Connect()

	for _, i := range dispatcher.Tasks {
		if i == nil {
			continue
		}

		response, err := i.Run(c)

		if err != nil {
			fmt.Println("Cannot run a task: " + i.Command)
		}

		fmt.Println(response)
	}

	fmt.Println("Done")
}
