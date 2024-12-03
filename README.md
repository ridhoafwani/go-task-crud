# Project task-crud

Simple CRUD application using gin framework

## Live Version

To make it easier for you to test the application, a live version is available at [https://task-crud.ridhoafwani.dev/swagger/index.html](https://task-crud.ridhoafwani.dev/swagger/index.html).

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

- Go (version 1.16 or later)
- Git

### Installation

1. Clone the repository

```bash
git clone https://github.com/ridhoafwani/go-task-crud.git
```

2. Install dependencies

```bash
cd task-crud
go mod tidy
```

3. Build the application

```bash
go build -o main cmd/main.go
```

or using Makefile

```bash
make build
```

4. Run the application

```bash
go run cmd/main.go
```

or using Makefile

```bash
make run
```

The application should now be running on `http://localhost:3000`.

## API Endpoints

The following endpoints are available:

- `GET /api/v1/tasks`: Get all tasks
- `POST /api/v1/tasks`: Create a new task
- `GET /api/v1/tasks/{id}`: Get a specific task
- `PATCH /api/v1/tasks/{id}`: Update a specific task
- `DELETE /api/v1/tasks/{id}`: Delete a specific task

For detailed API documentation, visit `http://localhost:3000/swagger/index.html` when the application is running.

## MakeFile Commands

Build the application

```bash
make build
```

Run the application

```bash
make run
```

Live reload the application:

```bash
make watch
```
