package config

type SSHConfig struct {
	Host    string
	Port    string
	Username string
	Password string
}

func NewSSHConfig(config *SSHConfig) *SSHConfig {
	if config.Port == "" {
		config.Port = "22"
	}

	if config.Host == "" {
		config.Host = "localhost"
	}

	if config.Username == "" {
		config.Username = "root"
	}

	if config.Password == "" {
		config.Password = ""
	}

	return config
}