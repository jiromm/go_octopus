### Octopus

Task based tool written in Go which allows you to play with remote server

```go
package main

import (
	o "./octopus"
)

func main() {
	dispatcher := &o.Dispatcher{
		Name: "Do some smart things",
		Config: o.NewConfig(&o.Config{
			Host: "0.0.0.0",
			Username: "root",
			Password: "toor",
		}),
	}

	dispatcher.AddTask(&o.Task{
		Name:    "Compress www directory",
		Command: "cd /var/ && zip -rq compressed.zip www/ && touch status.done",
		Type:	 o.TYPE_EXECUTE,
	})

	dispatcher.AddTask(&o.Task{
		Name:    "Make sure status.done exists",
		Command: "/var/status.done",
		Type:	 o.TYPE_EXISTENCE_CONFIDENCE,
	})

	dispatcher.AddTask(&o.Task{
		Name:    "Make sure compressed.zip exists",
		Command: "/var/compressed.zip",
		Type:	 o.TYPE_EXISTENCE_CONFIDENCE,
	})

	dispatcher.AddTask(&o.Task{
		Name:    "Remove status.done",
		Command: "/var/status.done",
		Type:	 o.TYPE_REMOVE,
	})

	dispatcher.AddTask(&o.Task{
		Name:    "Remove compressed.zip",
		Command: "/var/compressed.zip",
		Type:	 o.TYPE_REMOVE,
	})

	dispatcher.Run()
}
```