package octopus

func RunTestJob() {
	config := &Config{
		DB: NewDBConfig(&DBConfig{}),
		Filesystem: NewFilesystemConfig(&FilesystemConfig{}),
	}

	config.SSH = NewSSHConfig(&SSHConfig{
		Host: HOST,
		Username: USER,
		Password: PASS,
	})

	dispatcher := &Dispatcher{
		Name: "Do some smart things",
		Config: config,
	}

	dispatcher.AddTask(&Task{
		Name:    "Testing test task",
		Command: "test command",
		Type:	 TYPE_TEST,
	})

	dispatcher.Run()
}