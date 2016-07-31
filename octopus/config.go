package octopus

type Config struct {
	SSH *SSHConfig
	Filesystem *FilesystemConfig
	DB *DBConfig
}
