# Go To-Do

Welcome to the Go To-Do application! This project is a simple to-do list application built using Go. It allows users to manage their tasks via a RESTful API.

## Features

- **CRUD Operations**: Create, Read, Update, and Delete to-do items.
- **Persistent Storage**: Uses PostgreSQL for storing to-do items.
- **RESTful API**: Provides endpoints for managing tasks.

## Getting Started

### Prerequisites

- Go (1.18+)
- PostgreSQL

### Installation

1. **Clone the Repository**

   ```
   git clone https://github.com/obalaweb/go-todo.git
   cd go-todo
   ```

Set Up the Database

Ensure PostgreSQL is installed and running. Create a database for the application:

sql
Copy code
CREATE DATABASE todo_db;
Configuration

Create a .env file in the root directory and set the necessary environment variables:
```
DATABASE_URL=postgres://username:password@localhost:5432/todo_db
PORT=8080
```
Replace username and password with your PostgreSQL credentials.

Install Dependencies

Fetch the Go dependencies:
```
go mod tidy

```
Run Database Migrations

Apply database migrations to set up the required tables:
```
go run cmd/migrate/main.go
```
Build and Run the Application

Build the application:
```
go build -o go-todo main.go
```
Run the application:
```
./go-todo
```
The service will start on http://localhost:8080.

API Endpoints
GET /todos

Retrieve all to-do items.

Request:
```
GET /todos
```
Response:
```
[
  { "id": "1", "title": "Buy groceries", "completed": false },
  { "id": "2", "title": "Read a book", "completed": true }
]
POST /todos
```
Create a new to-do item.

Request:
```
POST /todos
Content-Type: application/json

{
  "title": "Learn Go",
  "completed": false
}
```
Response:
```
{
  "id": "3",
  "title": "Learn Go",
  "completed": false
}
```
```
GET /todos/{id}
```

Retrieve a specific to-do item by ID.

Request:
```
GET /todos/1
```
Response:
```
json
Copy code
{
  "id": "1",
  "title": "Buy groceries",
  "completed": false
}
```
```
PUT /todos/{id}
```
Update a specific to-do item by ID.

Request:
```
PUT /todos/1
Content-Type: application/json

{
  "title": "Buy groceries and snacks",
  "completed": true
}
```
Response:
```
{
  "id": "1",
  "title": "Buy groceries and snacks",
  "completed": true
}
DELETE /todos/{id}
```
Delete a specific to-do item by ID.

### Request:
```
DELETE /todos/1
```
Response:
```
{
  "status": "success"
}
```
### Configuration
Database Configuration: Update the DATABASE_URL in the .env file with your PostgreSQL credentials.
Port Configuration: Change the PORT variable in the .env file if needed.
Testing
Run tests using:
```
go test ./...
```
### Contributing
Contributions are welcome! Please open an issue or submit a pull request if you’d like to contribute.

License
This project is licensed under the MIT License - see the LICENSE file for details.

Contact
For any questions or feedback, please reach out to your.email@example.com.
