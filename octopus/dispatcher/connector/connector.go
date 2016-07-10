package connector

import (
	"fmt"

	c "../../config"
)

type Connector struct {
	Config *c.Config
}

func (connector *Connector) Connect() {
	fmt.Println("Connected to host: " + connector.Config.Host)
}