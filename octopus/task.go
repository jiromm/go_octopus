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

type Task struct {
	Name    string
	Command string
	Type	string
}

func (task *Task) Run(ssh *easyssh.MakeConfig) (result string, err error) {
	switch task.Type {
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
	fmt.Println("Executing '" + task.Command + "'")

	return ssh.Run(task.Command)
}

func (task *Task) CheckExistence(ssh *easyssh.MakeConfig) (result bool, err error) {
	fmt.Print("Checking existance of '" + task.Command + "'. ")

	r, e := ssh.Run("[ ! -e " + task.Command + " ]; echo $?")

	r = strings.Replace(r, "\n", "", -1)

	fmt.Println("Result is '" + r + "'")

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
	fmt.Println("Removing '" + task.Command + "'")

	_, err = ssh.Run("rm " + task.Command)

	return err
}

func (task *Task) Download(ssh *easyssh.MakeConfig) (err error) {
	fmt.Println("Downloading '" + task.Command + "'")

	err = ssh.Download(task.Command, "./storage/" + filepath.Base(task.Command))

	return err
}

func (task *Task) Upload(ssh *easyssh.MakeConfig) (err error) {
	fmt.Println("Uploading '" + task.Command + "'")

	err = ssh.Upload("./storage/" + task.Command, "~/" + filepath.Base(task.Command))

	return err
}
