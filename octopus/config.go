package octopus

import (
	c "./config"
)

type Config struct {
	SSH *c.SSHConfig
	Filesystem *c.FilesystemConfig
	DB *c.DBConfig
}
