package octopus

import (
	"github.com/jiromm/easyssh"
	"fmt"
)

type Connector struct {
	Config *Config
}

func (connector *Connector) Connect() (*easyssh.MakeConfig) {
	fmt.Println("Connecting to host: %s", connector.Config.SSH.Host)

	return &easyssh.MakeConfig{
		Server: connector.Config.SSH.Host,
		User: connector.Config.SSH.Username,
		Password: connector.Config.SSH.Password,
		Port: connector.Config.SSH.Port,
	}
}