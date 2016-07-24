package config

type DBConfig struct {
	Name string
}

func NewDBConfig(config *DBConfig) *DBConfig {
	if config.Name == "" {
		config.Name = "octopus"
	}

	return config
}