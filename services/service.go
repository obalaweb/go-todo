package services

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"
	"todo/db"
)

func NewTodoList() *TodoList {

	return &TodoList{
		todos: make(map[string]*Todo),
	}

}

type TodoList struct {
	todos map[string]*Todo
}

type Todo struct {
	Id          string
	Title       string
	Description string
	Completed   bool
	DueDate     string
	priority    Priority
	labels      []string
}

type Priority int

const (
	Lowest Priority = iota
	Low
	Medium
)

func (t *TodoList) AddTodo(title, description, dueDate string, priority Priority, labels []string) error {
	// Generate a new ID for the todo item
	id := GenerateID()

	// Convert the labels slice to a comma-separated string
	labelsStr := strings.Join(labels, ",")

	// Insert the new todo item into the database
	_, err := db.DB.ExecContext(
		context.Background(),
		`INSERT INTO todos (
			id,
			title, 
			description, 
			completed, 
			due_date, 
			priority, 
			labels
		) VALUES (?,?,?,?,?,?,?);`,
		id, title, description, false, dueDate, priority, labelsStr,
	)

	if err != nil {
		return err
	}

	// Create the new Todo object and store it in the map
	newTodo := &Todo{
		Id:          id,
		Title:       title,
		Description: description,
		Completed:   false,
		DueDate:     dueDate,
		priority:    priority,
		labels:      labels,
	}

	// Store the new todo in the TodoList
	t.todos[newTodo.Id] = newTodo

	return nil
}

func (t *TodoList) GetTodo(id string) (*Todo, error) {
	row := db.DB.QueryRow("SELECT id, title, description, completed, due_date, priority, labels FROM todos WHERE id = ?", id)
	todo := &Todo{}
	var labelsStr string
	err := row.Scan(&todo.Id, &todo.Title, &todo.Description, &todo.Completed, &todo.DueDate, &todo.priority, &labelsStr)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrTodoNotFound
		}
		return nil, err
	}
	todo.labels = strings.Split(labelsStr, ",")
	return todo, nil
}

func (t *TodoList) GetTodos() ([]*Todo, error) {
	rows, err := db.DB.Query("SELECT id, title, description, completed, due_date, priority, labels FROM todos")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	todos := []*Todo{}

	for rows.Next() {
		todo := &Todo{}

		var labelsStr string
		err := rows.Scan(&todo.Id, &todo.Title, &todo.Description, &todo.Completed, &todo.DueDate, &todo.priority, &labelsStr)
		if err == sql.ErrNoRows {
			return nil, ErrTodoNotFound
		}

		// Convert the comma-separated labels string back into a slice of strings
		todo.labels = strings.Split(labelsStr, ",")

		// Add the current todo to the slice
		todos = append(todos, todo)
	}

	// Check for errors after the loop
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return todos, nil
}

// todo, ok := t.todos[id]
// if !ok {
// 	return nil, ErrTodoNotFound
// }
// return todo, nil

func (t *TodoList) UpdateTodo(id string, title string, description string, dueDate string, completed bool) error {
	todo, err := t.GetTodo(id)
	if err != nil {
		return err
	}

	_, err = db.DB.Exec("UPDATE todos SET title =?, description =? , due_date =?, completed = ? WHERE id = ?",
		title, description, dueDate, completed, id,
	)

	if err != nil {
		return err
	}

	todo.Title = title
	todo.Description = description
	todo.DueDate = dueDate
	todo.Completed = completed
	return nil
}

func (t *TodoList) DeleteTodo(id string) error {
	_, err := t.GetTodo(id)
	if err != nil {
		return err
	}

	_, err = db.DB.Exec("DELETE FROM todos WHERE id = ?", id)

	if err != nil {
		return err
	}

	return nil
}

func (t *TodoList) CompleteTodo(id string) error {
	_, err := db.DB.Exec("UPDATE todos SET completed = ? WHERE id = ?", true, id)

	if err != nil {
		return err
	}

	todo, err := t.GetTodo(id)

	if err != nil {
		return err
	}
	todo.Completed = true
	return nil
}

func (t *TodoList) UncompleteTodo(id string) error {
	_, err := db.DB.Exec("UPDATE todos SET completed = ? WHERE id = ?", false, id)

	if err != nil {
		return err
	}

	todo, err := t.GetTodo(id)

	if err != nil {
		return err
	}

	todo.Completed = false
	return nil
}

func (t *TodoList) GetTodosByLabel(label string) ([]*Todo, error) {
	var matchingTodos []*Todo
	for _, todo := range t.todos {
		for _, l := range todo.labels {
			if l == label {
				matchingTodos = append(matchingTodos, todo)
				break
			}
		}
	}
	if len(matchingTodos) == 0 {
		return nil, ErrNoTodosWithLabel
	}
	return matchingTodos, nil
}

func (t *TodoList) GetTodosByPriority(priority Priority) ([]*Todo, error) {
	var matchingTodos []*Todo
	for _, todo := range t.todos {
		if todo.priority == priority {
			matchingTodos = append(matchingTodos, todo)
		}
	}
	if len(matchingTodos) == 0 {
		return nil, ErrNoTodosWithPriority
	}
	return matchingTodos, nil
}

func (t *TodoList) GetTodosByDueDate(dueDate string) ([]*Todo, error) {
	var matchingTodos []*Todo
	for _, todo := range t.todos {
		if todo.DueDate == dueDate {
			matchingTodos = append(matchingTodos, todo)
		}
	}
	if len(matchingTodos) == 0 {
		return nil, ErrNoTodosWithDueDate
	}
	return matchingTodos, nil
}

func (t *TodoList) GetTodosByStatus(completed bool) ([]*Todo, error) {
	var matchingTodos []*Todo
	for _, todo := range t.todos {
		if todo.Completed == completed {
			matchingTodos = append(matchingTodos, todo)
		}
	}
	if len(matchingTodos) == 0 {
		return nil, ErrNoTodosWithStatus
	}
	return matchingTodos, nil
}

func (t *TodoList) GetTodosByTitle(title string) ([]*Todo, error) {
	var matchingTodos []*Todo
	for _, todo := range t.todos {
		if todo.Title == title {
			matchingTodos = append(matchingTodos, todo)
		}
	}
	if len(matchingTodos) == 0 {
		return nil, ErrNoTodosWithTitle
	}
	return matchingTodos, nil
}

func (t *TodoList) GetTodosByDescription(description string) ([]*Todo, error) {
	var matchingTodos []*Todo
	for _, todo := range t.todos {
		if todo.Description == description {
			matchingTodos = append(matchingTodos, todo)
		}
	}
	if len(matchingTodos) == 0 {
		return nil, ErrNoTodosWithDescription
	}
	return matchingTodos, nil
}

func GenerateID() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}

func GetCurrentTime() string {
	return time.Now().Format(time.RFC3339)
}

var (
	ErrTodoNotFound           = errors.New("todo not found")
	ErrNoTodosWithLabel       = errors.New("no todos found with the given label")
	ErrNoTodosWithPriority    = errors.New("no todos found with the given priority")
	ErrNoTodosWithDueDate     = errors.New("no todos found with the given due date")
	ErrNoTodosWithStatus      = errors.New("no todos found with the given status")
	ErrNoTodosWithAuthor      = errors.New("no todos found with the given author")
	ErrNoTodosWithTitle       = errors.New("no todos found with the given title")
	ErrNoTodosWithDescription = errors.New("no todos found with the given description")
	ErrTodoTitleExists        = errors.New("a todo with the given title already exists")
	ErrTodoDescriptionExists  = errors.New("a todo with the given description already exists")
	ErrTodoIDExists           = errors.New("a todo with the given ID already exists")
	ErrTodoTitleTooLong       = errors.New("todo title is too long")
	ErrTodoDescriptionTooLong = errors.New("todo description is too long")
	ErrTodoDueDateInvalid     = errors.New("invalid due date format")
	ErrTodoPriorityInvalid    = errors.New("invalid priority value")
	ErrTodoLabelTooLong       = errors.New("todo label is too long")
	ErrTodoCommentTextTooLong = errors.New("comment text is too long")
	ErrTodoParentIDInvalid    = errors.New("invalid parent ID")
)
