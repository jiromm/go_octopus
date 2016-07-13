package main

func main() {
	dispatcher := &Dispatcher{
		Name: "Migrate alpha's database",
		Config: NewConfig(&Config{}),
	}

	task1 := &Task{
		Name:    "Find php files in www folder",
		Command: "cd /var/www && ls -la | grep php",
	}

	task2 := &Task{
		Name:    "Restart a server",
		Command: "reboot",
	}

	dispatcher.AddTask(task1)
	dispatcher.AddTask(task2)

	dispatcher.Run()
}
