package octopus

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func BuildEnvironment(db *sql.DB) (err error) {
	queries := ReturnQueries()

	err = CreateTables(db, queries)

	if err != nil {
		return err
	}

	return nil
}

func ReturnQueries() (queries []string) {
	queries = []string{
		`CREATE TABLE IF NOT EXISTS session (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name VARCHAR(128) NULL,
			status CHAR(1) NOT NULL DEFAULT ('P'), -- P - pending, O - ongoing, F - finished
			started_at DATE NULL,
			ended_at DATE NULL
		);`,
		`CREATE TABLE IF NOT EXISTS task (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			session_id INTEGER,
			name VARCHAR(128) NULL,
			status CHAR(1) NOT NULL DEFAULT ('P'), -- P - pending, O - ongoing, F - finished
			started_at DATE NULL,
			ended_at DATE NULL,

			FOREIGN KEY (session_id) REFERENCES session(id)
		);`,
	}

	return queries
}

func CreateTables(db *sql.DB, queries []string) (err error) {
	for _, value := range queries {
		_, err := db.Exec(value)

		if err != nil {
			return err
		}
	}

	return nil
}