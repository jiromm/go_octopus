package main

import (
	s "./session"
	t "./task"
	c "./config"
)

func main() {
	session := &s.Session{
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

	session.AddTask(task1)
	session.AddTask(task2)

	session.Run()
}
