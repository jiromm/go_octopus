package octopus

import (
	"fmt"
	"time"
	"strings"
	"path/filepath"

	"github.com/jiromm/easyssh"
)

const TYPE_EXECUTE = "execute"
const TYPE_EXISTENCE_CONFIDENCE = "existence_confidence"
const TYPE_REMOVE = "remove"
const TYPE_DOWNLOAD = "download"
const TYPE_UPLOAD = "upload"
const TYPE_TEST = "test"

type Task struct {
	Name    string
	Command string
	Type	string
	UUId	int64
	Config	*Config
}

func (task *Task) SetConfig(config *Config) {
	task.Config = config
}

func (task *Task) SetUUId(uuid int64) {
	task.UUId = uuid
}

func (task *Task) Run(ssh *easyssh.MakeConfig) (result string, err error) {
	switch task.Type {
	case TYPE_TEST:
		err = nil
	case TYPE_EXECUTE:
		result, err = task.Execute(ssh)
	case TYPE_EXISTENCE_CONFIDENCE:
		is := false

		for !is {
			is, err = task.CheckExistence(ssh)

			if err != nil {
				panic("Something went wrong on existence check")
			}

			time.Sleep(2 * time.Second)
		}
	case TYPE_REMOVE:
		err = task.Remove(ssh)
	case TYPE_DOWNLOAD:
		err = task.Download(ssh)
	case TYPE_UPLOAD:
		err = task.Upload(ssh)
	}

	return result, err
}

func (task *Task) Execute(ssh *easyssh.MakeConfig) (outStr string, err error) {
	fmt.Println("Executing '%s'", task.Command)

	return ssh.Run(task.Command)
}

func (task *Task) CheckExistence(ssh *easyssh.MakeConfig) (result bool, err error) {
	fmt.Print("Checking existance of '%s'", task.Command)

	r, e := ssh.Run(fmt.Sprintf("[ ! -e %s ]; echo $?", task.Command))

	r = strings.Replace(r, "\n", "", -1)

	fmt.Println("Result is '%s'", r)

	if e != nil {
		err = e
	}

	if r == "1" {
		result = true
	} else {
		result = false
	}

	return result, nil
}

func (task *Task) Remove(ssh *easyssh.MakeConfig) (err error) {
	fmt.Println("Removing '%s'", task.Command)

	_, err = ssh.Run(fmt.Sprintf("rm %s", task.Command))

	return err
}

func (task *Task) Download(ssh *easyssh.MakeConfig) (err error) {
	fmt.Println("Downloading '%s'", task.Command)

	err = ssh.Download(
		task.Command,
		fmt.Sprintf("./%s/%s/%s",
			task.Config.Filesystem.Storage,
			task.Config.Filesystem.Files,
			filepath.Base(task.Command)))

	return err
}

func (task *Task) Upload(ssh *easyssh.MakeConfig) (err error) {
	fmt.Println("Uploading '%s'", task.Command)

	err = ssh.Upload(
		fmt.Sprintf("./%s/%s//%s",
			task.Config.Filesystem.Storage,
			task.Config.Filesystem.Files,
			task.Command),
		fmt.Sprintf("~/%s", filepath.Base(task.Command)))

	return err
}
