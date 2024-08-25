package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"todo/db"
	"todo/services"

	"github.com/markkurossi/tabulate"
)

func main() {
	db.Connector()
	// rowId := 1724521252616378149

startOver:
	choice := 0
	fmt.Printf("What do you want to do?\n 1. List all todos\n 2. Add new Todo\n 3. Get Todo by ID\n 4. Update Todo\n 5. Delete a Todo\n 6. Complete a Todo \n")
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
	var title, description, dueDate string
	var priority services.Priority
	var labels []string

	fmt.Print("Enter Todo Title: ")
	fmt.Scanln(&title)

	fmt.Println("Enter Todo Description: ")
	result := make([]byte, 1024)
	in := bufio.NewReader(os.Stdin)
	_, errs := in.Read(result)
	if errs != nil {
		log.Fatal(errs)
		return
	}

	description = string(result)

	fmt.Print("Enter Todo Due Date (YYYY-MM-DD): ")
	fmt.Scanln(&dueDate)

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

	var title, description, dueDate string

	fmt.Print("Enter Todo Title (leave blank to keep current): ")
	fmt.Scanln(&title)
	if title != "" {
		todo.Title = title
	}

	fmt.Print("Enter Todo Description (leave blank to keep current): ")
	fmt.Scanln(&description)
	if description != "" {
		todo.Description = description
	}

	fmt.Print("Enter Todo Due Date (YYYY-MM-DD) (leave blank to keep current): ")
	fmt.Scanln(&dueDate)
	if dueDate != "" {
		todo.DueDate = dueDate
	}
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
