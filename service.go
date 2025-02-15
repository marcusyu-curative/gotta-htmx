package main

import (
	"database/sql"
	"log"

	"examples/htmx-go/models"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDatabase() {
	var err error
	DB, err = sql.Open("sqlite3", "./awesome.db")
	if err != nil {
		log.Fatal(err)
	}

	_, err = DB.Exec(`
	CREATE TABLE IF NOT EXISTS todos (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT,
		status TEXT
	);`)

	if err != nil {
		log.Fatal(err)
	}
}

func CreateToDo(title string, status string) (int64, error) {
	result, err := DB.Exec("INSERT INTO todos (title, status) VALUES (?, ?)", title, status)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func DeleteToDo(id int64) error {
	_, err := DB.Exec("DELETE FROM todos WHERE id = ?", id)
	return err
}

func ReadToDoList() []models.ToDo {
	rows, _ := DB.Query("SELECT id, title, status FROM todos")
	defer rows.Close()

	todos := make([]models.ToDo, 0)
	for rows.Next() {
		var todo models.ToDo
		rows.Scan(&todo.Id, &todo.Title, &todo.Status)
		todos = append(todos, todo)
	}
	return todos
}
