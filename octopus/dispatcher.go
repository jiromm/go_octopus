package octopus

import (
	"fmt"

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
	sshConnector := SSHConnector{
		Config: dispatcher.Config,
	}
	dbConnector := DBConnector{
		Config: dispatcher.Config,
	}

	ssh := sshConnector.Connect()
	db := dbConnector.Connect()

	err := BuildEnvironment(db)

	if err != nil {
		panic(err)
	}

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

		if response != "" {
			fmt.Println(response)
		}
	}

	fmt.Println("Done")
}

func (dispatcher *Dispatcher) LetItStart(db *sql.DB) int64 {
	sessionId, err := dispatcher.CreateSession(db)

	if err != nil {
		panic(err)
	}

	for _, task := range dispatcher.Tasks {
		if task == nil {
			continue
		}

		taskId, err := dispatcher.CreateTask(db, task, sessionId)

		if err != nil {
			panic(err)
		}

		task.SetUUId(taskId)
	}

	return sessionId
}

func (dispatcher *Dispatcher) CreateSession(db *sql.DB) (int64, error) {
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

	fmt.Println("Session has been created");

	return res.LastInsertId()
}

func (dispatcher *Dispatcher) CreateTask(db *sql.DB, task *Task, sessionId int64) (int64, error) {
	stmt, err := db.Prepare(`
		INSERT INTO task (
			session_id,
			name,
			status,
			started_at
		) values (?, ?, ?, ?);
	`)

	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	res, err2 := stmt.Exec(sessionId, task.Name, "P", time.Now())

	if err2 != nil {
		panic(err2)
	}

	fmt.Println("Task has been created: ", task.Name);

	return res.LastInsertId()
}

func (dispatcher *Dispatcher) MarkTaskAsDone(db *sql.DB, task *Task) {
	stmt, err := db.Prepare(`
		UPDATE task
		SET status = ?, ended_at = ?
		WHERE id = ?;
	`)

	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	_, err2 := stmt.Exec("P", time.Now(), task.UUId)

	if err2 != nil {
		panic(err2)
	}

	fmt.Println("Task '%s' marked as done", task.Name);
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

	fmt.Println("Session with id %d marked as done", sessionId);
}

func (dispatcher *Dispatcher) LetItEnd(db *sql.DB, sessionId int64) {
	dispatcher.CloseSession(db, sessionId)
	db.Close()

	fmt.Println("DB Connection has been closed")
}
