package octopus

import (
	"fmt"

	c "./connector"
	"database/sql"
	"time"
)

var I = 0

type Dispatcher struct {
	Name  string
	Tasks [10]*Task
	Config *Config
}

func (dispatcher *Dispatcher) AddTask(task *Task) {
	task.SetConfig(dispatcher.Config)

	dispatcher.Tasks[I] = task
	I += 1
}

func (dispatcher *Dispatcher) Run() {
	sshConnector := c.SSHConnector{
		Config: dispatcher.Config,
	}
	dbConnector := c.DBConnector{
		Config: dispatcher.Config,
	}

	ssh := sshConnector.Connect()
	db := dbConnector.Connect()

	sessionId := dispatcher.LetItStart(db)
	defer dispatcher.LetItEnd(db, sessionId)

	for _, task := range dispatcher.Tasks {
		if task == nil {
			continue
		}

		response, err := task.Run(ssh)

		if err != nil {
			fmt.Println("Cannot run a task: %s", task.Command)
		}

		dispatcher.MarkTaskAsDone(db, task)

		fmt.Println(response)
	}

	fmt.Println("Done")
}

func (dispatcher *Dispatcher) LetItStart(db *sql.DB) int64 {
	sessionId := dispatcher.CreateSession(db)

	for _, task := range dispatcher.Tasks {
		taskId := dispatcher.CreateTask(db, task, sessionId)
		task.SetUUId(taskId)
	}

	return sessionId
}

func (dispatcher *Dispatcher) CreateSession(db *sql.DB) int64 {
	stmt, err := db.Prepare(`
		INSERT INTO session (
			name,
			status,
			started_at
		) values (?, ?, ?)
	`)

	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	res, err2 := stmt.Exec(dispatcher.Name, "O", time.Now())

	if err2 != nil {
		panic(err2)
	}

	return res.LastInsertId()
}

func (dispatcher *Dispatcher) CreateTask(db *sql.DB, task *Task, sessionId int64) int64 {
	stmt, err := db.Prepare(`
		INSERT INTO task (
			session_id,
			name,
			status,
			started_at
		) values (?, ?, ?, ?)
	`)

	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	res, err2 := stmt.Exec(sessionId, task.Name, "P", time.Now())

	if err2 != nil {
		panic(err2)
	}

	return res.LastInsertId()
}

func (dispatcher *Dispatcher) MarkTaskAsDone(db *sql.DB, task *Task) {
	stmt, err := db.Prepare(`
		UPDATE task
		SET status = ?, ended_at = ?
		WHERE id = ?
	`)

	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	_, err2 := stmt.Exec("P", time.Now(), task.UUId)

	if err2 != nil {
		panic(err2)
	}
}

func (dispatcher *Dispatcher) CloseSession(db *sql.DB, sessionId int64) {
	stmt, err := db.Prepare(`
		UPDATE task
		SET status = ?, ended_at = ?
		WHERE id = ?
	`)

	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	_, err2 := stmt.Exec("F", time.Now(), sessionId)

	if err2 != nil {
		panic(err2)
	}
}

func (dispatcher *Dispatcher) LetItEnd(db *sql.DB, sessionId int64) {
	fmt.Println("DB Connection has been closed")

	dispatcher.CloseSession(db, sessionId)
	db.Close()
}
