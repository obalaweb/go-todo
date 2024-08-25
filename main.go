package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"todo/db"
	"todo/services"

	"github.com/markkurossi/tabulate"
)

func main() {
	db.Connector()
	// rowId := 1724521252616378149

startOver:
	choice := 0
	fmt.Printf("What do you want to do?\n 1. List all Todos\n 2. Add new Todo\n 3. Get Todo by ID\n 4. Update Todo\n 5. Delete a Todo\n 6. Complete a Todo \n")
	fmt.Scanln(&choice)

	if choice == 0 {
		fmt.Println("Exiting...")
		return
	}

	switch choice {
	case 1:
		listTodos()
		goto startOver
	case 2:
		addTodo()
		goto startOver

	case 3:
		id := ""
		fmt.Printf("Enter Todo ID: ")
		fmt.Scanln(&id)
		getTodoById(id)
		goto startOver

	case 4:
		id := ""
		fmt.Printf("Enter Todo ID: ")
		fmt.Scanln(&id)
		updateTodo(id)
		goto startOver

	case 5:
		id := ""
		fmt.Printf("Enter Todo ID: ")
		fmt.Scanln(&id)
		deleteTodo(id)
		goto startOver

	case 6:
		id := ""
		fmt.Printf("Enter Todo ID: ")
		fmt.Scanln(&id)
		completeTodo(id)
		goto startOver

	}

}

func listTodos() {
	todoLists, err := services.NewTodoList().GetTodos()

	if err != nil {
		log.Fatal(err)
	}

	tab := tabulate.New(tabulate.Unicode)
	tab.Header("id")
	tab.Header("Title").SetAlign(tabulate.ML)
	tab.Header("Description").SetAlign(tabulate.ML)
	tab.Header("Completed").SetAlign(tabulate.ML)
	for _, todo := range todoLists {
		row := tab.Row()
		row.Column(todo.Id)
		row.Column(todo.Title)
		row.Column(todo.Description)
		row.Column(fmt.Sprintf("%v", todo.Completed))
	}

	tab.Print(os.Stdout)
}

func addTodo() {
	var priority services.Priority
	var labels []string
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Todo Title: ")
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)

	fmt.Println("Enter Todo Description: ")
	description, _ := reader.ReadString('\n')
	description = strings.TrimSpace(description)

	fmt.Print("Enter Todo Due Date (YYYY-MM-DD): ")
	dueDate, _ := reader.ReadString('\n')
	dueDate = strings.TrimSpace(dueDate)

	fmt.Print("Enter Todo Priority (0: Lowest, 1: Low, 2: Medium): ")
	fmt.Scanln(&priority)

	fmt.Print("Enter Todo Labels (comma-separated): ")
	fmt.Scanln(&labels)

	err := services.NewTodoList().AddTodo(title, description, dueDate, priority, labels)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Todo added successfully")
}

func getTodoById(id string) {
	todo, err := services.NewTodoList().GetTodo(id)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Title: %s\nDescription: %s\nDue Date: %s\n", todo.Title, todo.Description, todo.DueDate)
}

func updateTodo(id string) {
	todo, err := services.NewTodoList().GetTodo(id)

	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter new title: ")
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)
	if title == "" {
		title = todo.Title
	}

	fmt.Print("Enter new description: ")
	description, _ := reader.ReadString('\n')
	description = strings.TrimSpace(description)
	if description == "" {
		description = todo.Description
	}

	fmt.Print("Enter new due date (YYYY-MM-DD): ")
	dueDate, _ := reader.ReadString('\n')
	dueDate = strings.TrimSpace(dueDate)
	if dueDate == "" {
		dueDate = todo.DueDate
	}

	fmt.Print("Enter completed status (true/false): ")
	completedStr, _ := reader.ReadString('\n')
	completedStr = strings.TrimSpace(completedStr)
	completed, err := strconv.ParseBool(completedStr)
	if err != nil {
		fmt.Println("Invalid completed status.")
		return
	}

	err = services.NewTodoList().UpdateTodo(id, title, description, dueDate, completed)
	if err != nil {
		fmt.Println("Error updating todo:", err)
		return
	}

	fmt.Println("Todo updated successfully!")
}

func deleteTodo(id string) {
	err := services.NewTodoList().DeleteTodo(id)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Todo deleted successfully")
}

func completeTodo(id string) {
	err := services.NewTodoList().CompleteTodo(id)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Todo completed successfully")
}
