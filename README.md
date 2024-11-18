# CRUD API with Go and MySQL

This repository demonstrates a **CRUD API** built using **Go (Golang)** with **MySQL** and includes foundational concepts such as routing, request handling, and database interactions.

The project is structured to separate concerns, making it easier to maintain and scale.


## Project Structure
```
cmd/
  └── main/                # Contains the main.go file to run the application
pkg/
  ├── config/              # Database configuration (app.go)
  ├── controllers/         # Controller for processing requests and responses
  ├── models/              # Database model and struct (book.go)
  ├── routes/              # Routes for API endpoints
  └── utils/               # Utility functions for JSON marshalling/unmarshalling

```


## Features
+ Create, Read, Update, and Delete operations for book records.
+ RESTful routing using Gorilla Mux.
+ Seamless integration with MySQL using GORM ORM.
+ Clean and modular code structure.


## Prerequisites
+ Go (1.16 or higher)
+ MySQL installed and running
+ VS Code or any IDE for Go development


## Setup Instructions

### 1. Clone the Repository
```
git clone https://github.com/Chinzzii/CRUD-API-with-MySQL-and-Go.git
cd CRUD-API-with-MySQL-and-Go
```

### 2. Install Dpendencies
Install the necessary Go packages:
```
go mod tidy
```

### 3. Configure the Database
+ Set up a MySQL database.
+ Update the ```config/app.go``` file with your database credentials.

### 4. Run the Application
Start the API server:
```
go run cmd/main/main.go
```
The server will run at ```http://localhost:9010```.


## API Endpoints
| Method | Endpoint        | Description                     |
|--------|-----------------|---------------------------------|
| GET    | `/books`        | Retrieve all books              |
| POST   | `/books`        | Add a new book                  |
| GET    | `/books/{id}`   | Retrieve a book by ID           |
| PUT    | `/books/{id}`   | Update a book by ID             |
| DELETE | `/books/{id}`   | Delete a book by ID             |


## Dependencies
The project uses the following Go modules:
+ GORM: For interacting with MySQL.
```
go get github.com/jinzhu/gorm
go get github.com/jinzhu/gorm/dialects/mysql
```
+ Gorilla Mux: For routing.
```
go get github.com/gorilla/mux
```


## Contribution
Contributions are welcome! Feel free to fork the repo and open a pull request with your changes.
