package todo

import (
	"fmt"
	"log"

	"github.com/deepkush97/go-todo-sqlite/internal/db"
)

type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func fetchAllTodos() ([]Todo, error) {
	rows, err := db.GetDB().Query("SELECT id, title, completed FROM todos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []Todo
	for rows.Next() {
		var todo Todo
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Completed); err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	return todos, nil
}

func getOneTodo(id int) (*Todo, error) {

	stmt, err := db.GetDB().Prepare("SELECT id, title, completed FROM todos where id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}

	todo := Todo{}

	defer rows.Close()

	found := false

	for rows.Next() {
		found = true
		err = rows.Scan(&todo.ID, &todo.Title, &todo.Completed)
		if err != nil {
			log.Printf("row scan error: %v", err)
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	if !found {
		return nil, fmt.Errorf("not found")
	}

	return &todo, nil
}

func insertTodo(title string) error {
	query := "INSERT INTO todos (title) VALUES (?)"
	_, err := db.GetDB().Exec(query, title)
	return err
}

func updateTodo(id int, title string) error {
	query := "UPDATE todos set title=? WHERE id=?"
	_, err := db.GetDB().Exec(query, title, id)
	return err
}

func markTodoCompleted(id int) error {
	query := "UPDATE todos SET completed = TRUE WHERE id = ?"
	_, err := db.GetDB().Exec(query, id)
	return err
}

func deleteTodoByID(id int) error {
	query := "DELETE FROM todos WHERE id = ?"
	_, err := db.GetDB().Exec(query, id)
	return err
}
