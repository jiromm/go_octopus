package octopus

import (
	"github.com/jiromm/easyssh"
	"fmt"
)

type SSHConnector struct {
	Config *Config
}

func (connector *SSHConnector) Connect() (*easyssh.MakeConfig) {
	fmt.Println("Connecting to host: ", connector.Config.SSH.Host)

	return &easyssh.MakeConfig{
		Server: connector.Config.SSH.Host,
		User: connector.Config.SSH.Username,
		Password: connector.Config.SSH.Password,
		Port: connector.Config.SSH.Port,
	}
}