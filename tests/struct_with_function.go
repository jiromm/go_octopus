package main

import (
	t "./task"
)

func main() {
	task := &t.Task{
		Name:    "Find php files in www folder",
		Command: "cd /var/www && ls -la",
	}

	task.Execute()
}
