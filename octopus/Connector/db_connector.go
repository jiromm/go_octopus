package connector

import (
	"fmt"

	c "../"
	"database/sql"
)

type DBConnector struct {
	Config *c.Config
}

func (connector *DBConnector) Connect() (*sql.DB) {
	fmt.Println("Connecting to db: %s", connector.Config.DB.Name)

	db, err := sql.Open("sqlite3", fmt.Sprintf("../%s/%s", connector.Config.Filesystem.Storage, connector.Config.DB.Dir))

	if err != nil {
		panic(err)
	}

	if db == nil {
		panic("db nil")
	}

	return db
}