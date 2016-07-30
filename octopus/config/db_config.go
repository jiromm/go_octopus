package config

type DBConfig struct {
	Name string
	Dir string
}

func NewDBConfig(config *DBConfig) *DBConfig {
	if config.Name == "" {
		config.Name = "octopus"
	}

	if config.Dir == "" {
		config.Dir = "db/session.db"
	}

	return config
}