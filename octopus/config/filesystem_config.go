package config

type FilesystemConfig struct {
	Storage string
	Files string
}

func NewFilesystemConfig(config *FilesystemConfig) *FilesystemConfig {
	if config.Storage == "" {
		config.Storage = "storage"
	}

	if config.Files == "" {
		config.Files = "files"
	}

	return config
}