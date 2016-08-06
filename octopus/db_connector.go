package octopus

import (
	"fmt"
	"database/sql"
)

type DBConnector struct {
	Config *Config
}

func (connector *DBConnector) Connect() (*sql.DB) {
	fmt.Printf("Connecting to db [%s]\n", connector.Config.DB.Name)

	db, err := sql.Open("sqlite3", fmt.Sprintf("./%s/%s", connector.Config.Filesystem.Storage, connector.Config.DB.Dir))

	if err != nil {
		panic(err)
	}

	if db == nil {
		panic("db nil")
	}

	return db
}