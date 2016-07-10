package config

type Config struct {
	Host    string
	Port    int
	Username string
	Password string
}

func NewConfig(config *Config) *Config {
	if config.Port == 0 {
		config.Port = 22
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