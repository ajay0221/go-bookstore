# Go Bookstore

This is a simple RESTful API for a bookstore, built with Go.

## Technologies Used

- Go
- Gorilla Mux - A powerful HTTP router and URL matcher for building Go web servers.
- Gorm - The fantastic ORM library for Golang, aims to be developer friendly.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

- Go (at least version 1.16)
- MySQL

### Installing

1. Clone the repository
`git clone https://github.com/ajay0221/go-bookstore.git`
2. Navigate to the project directory
`cd go-bookstore`
3. Install the dependencies
`go mod download`
4. Update the database connection details in `pkg/config/config.go`
5. Run the server
`go run main.go`
6. The server will start on `localhost:9010`.

## API Endpoints
| Method | Endpoint     | Description                           |
|--------|--------------|---------------------------------------|
| GET    | /books       | Retrieves a list of all books         |
| GET    | /books/{id}  | Retrieves a specific book by its ID   |
| POST   | /books       | Creates a new book                    |
| PUT    | /books/{id}  | Updates a specific book by its ID     |
| DELETE | /books/{id}  | Deletes a specific book by its ID     |