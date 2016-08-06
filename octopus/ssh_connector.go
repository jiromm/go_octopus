package octopus

import (
	"github.com/jiromm/easyssh"
	"fmt"
)

type SSHConnector struct {
	Config *Config
}

func (connector *SSHConnector) Connect() (*easyssh.MakeConfig) {
	fmt.Printf("Connecting to host [%s]\n", connector.Config.SSH.Host)

	return &easyssh.MakeConfig{
		Server: connector.Config.SSH.Host,
		User: connector.Config.SSH.Username,
		Password: connector.Config.SSH.Password,
		Port: connector.Config.SSH.Port,
	}
}