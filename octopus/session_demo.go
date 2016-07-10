package main

import (
	d "./dispatcher"
	t "./task"
	c "./config"
)

func main() {
	dispatcher := &d.Dispatcher{
		Name: "Migrate alpha's database",
		Config: c.NewConfig(&c.Config{}),
	}

	task1 := &t.Task{
		Name:    "Find php files in www folder",
		Command: "cd /var/www && ls -la | grep php",
	}

	task2 := &t.Task{
		Name:    "Restart a server",
		Command: "reboot",
	}

	dispatcher.AddTask(task1)
	dispatcher.AddTask(task2)

	dispatcher.Run()
}
