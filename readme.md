# Final Project 01
![Project Status](https://img.shields.io/badge/status-in%20progress-yellow)
![Contributors](https://img.shields.io/badge/contributors-3-blue)
![License](https://img.shields.io/badge/license-MIT-green)

A simple RESTful API for managing to-do tasks. This project is part of the Final Project 01 for the Magang dan Studi Independen Bersertifikat (MSIB) program at Hacktiv8.

## Team Members

- Ahmad Arsy (GLNG-KS-08-05)
- Eki Alfani (GLNG-KS08-012)
- Fabianus Jericho Harapan Jaya Sipayung (GLNG-KS08-015)

## Project Overview

This project is a RESTful API that allows users to perform basic CRUD (Create, Read, Update, Delete) operations on to-do tasks. It is designed as part of the Final Project 01 for the Magang dan Studi Independen Bersertifikat (MSIB) program at Hacktiv8.

## Features

- Create a new to-do task.
- Retrieve a list of all to-do tasks.
- Retrieve details of a specific to-do task by ID.
- Update information of an existing to-do task.
- Delete a to-do task.

## Technologies

The project is built using the following technologies:

- Programming Language: Golang
- Framework: Gin-Gonic
- Database: PostgreSQL
- HTTP Server: Gin HTTP Server
- API Docs: Swagger

## Installation
1. Clone the repo
```sh
  git clone https://github.com/hacktiv8-fp-golang/final-project-01.git
```
2. Install the required dependencies
```sh
  go get
```
3. Navigate to the root directory
```sh
  cd final-project-01
```
4. Create a .env file in the root directory
```sh
  touch .env
  # or
  echo > .env
```
5. Add environment variable
```sh
  DB_HOST=Database host (example: localhost)
  DB_USER=Database username (example: postgres)
  DB_PASSWORD=Database user's password (example: postgres)
  DB_PORT=Database port (example: 5432)
  DB_NAME=Name of the database to be used (example: postgres)
  DB_SSLMODE=Database SSL mode (example: disable)
  PORT=Port for the server to run on (example: 8080)
```
6. Run the project
```sh
  go run main.go
```
7. The project will be available at http://localhost:8080 by default.

## Usage

You can use various API endpoints to interact with the to-do tasks. See the API Endpoints section for details.

## API Endpoints

- **GET /todos**: Retrieve a list of all to-do tasks.

- **GET /todos/{id}**: Retrieve a specific to-do task by ID.

- **POST /todos**: Create a new to-do task.

- **PUT /todos/{id}**: Update an existing to-do task by ID.
- **DELETE /todos/{id}**: Delete a to-do task by ID.

## Examples

### Creating a New To-Do Task

#### Request:

```http
  POST /todos
  Content-Type: application/json

  {
      "title": "Buy Groceries",
      "completed": false
  }
```

#### Response (201 Created):
```json
  {
    "id": 1,
    "title": "Buy Groceries",
    "completed": false,
    "created_at": "2023-10-30T19:55:52.32604747+08:00",
    "updated_at": "2023-10-30T19:55:52.32604747+08:00"
  }
```

### Getting a List of To-Do Tasks

#### Request:

```http
  GET /todos
```

#### Response (200 OK):
```json
  [
    {
      "id": 1,
      "title": "Buy Groceries",
      "completed": false,
      "created_at": "2023-10-30T19:55:52.32604747+08:00",
      "updated_at": "2023-10-30T19:55:52.32604747+08:00"
    }
  ]
```

### Getting Details of a To-Do Task

#### Request:

```http
  GET /todos/1
```

#### Response (200 OK):
```json
  {
    "id": 1,
    "title": "Buy Groceries",
    "completed": false,
    "created_at": "2023-10-30T19:55:52.32604747+08:00",
    "updated_at": "2023-10-30T19:55:52.32604747+08:00"
  }
```
### Updating a To-Do Task

#### Request:

```http
  PUT /todos/1
  Content-Type: application/json

  {
      "title": "Buy Groceries",
      "completed": true
  }
```

#### Response (200 OK):
```json
  {
    "id": 1,
    "title": "Buy Groceries",
    "completed": true,
    "created_at": "2023-10-30T19:55:52.32604747+08:00",
    "updated_at": "2023-10-30T19:56:30.32604747+08:00"
  }
```

### Deleting a To-Do Task

#### Request:

```http
  DELETE /todos/1
```

#### Response (200 OK):
```json
  {
    "message": "Todo with id 1 has been successfully deleted"
  }
```

## API Documentation
The API documentation can be found [here](https://final-project-01-production.up.railway.app/swagger/index.html).

## Deployment
The API has been deployed and can be accessed [here](https://final-project-01-production.up.railway.app/).

## Contributing

We welcome contributions from developers to improve this project. Please check the [ISSUES](https://github.com/hacktiv8-fp-golang/final-project-01/issues) for a list of tasks that need to be done.

## License

This project is licensed under the [MIT License](LICENSE).
