package octopus

import (
	"github.com/jiromm/easyssh"
	"fmt"
)

type Connector struct {
	Config *Config
}

func (connector *Connector) Connect() (*easyssh.MakeConfig) {
	fmt.Println("Connecting to host: " + connector.Config.Host)

	return &easyssh.MakeConfig{
		Server: connector.Config.Host,
		User: connector.Config.Username,
		Password: connector.Config.Password,
		Port: connector.Config.Port,
	}
}