package models

import "database/sql"

type ToDo struct {
	Id     int64  `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

type ToDoModel struct {
    DB *sql.DB
}

func (t ToDoModel) Create(title string, status string) (int64, error) {
    result, err := t.DB.Exec(`
        INSERT INTO todos (title, status) VALUES (?, ?)`, title, status)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (t ToDoModel) Delete(id int64) error {
    _, err := t.DB.Exec("DELETE FROM todos WHERE id = ?", id)
	return err
}

func (t ToDoModel) ReadToDoList() []ToDo {
	rows, _ := t.DB.Query("SELECT id, title, status FROM todos")
	defer rows.Close()

	todos := make([]ToDo, 0)
	for rows.Next() {
		var todo ToDo
		rows.Scan(&todo.Id, &todo.Title, &todo.Status)
		todos = append(todos, todo)
	}
	return todos
}
