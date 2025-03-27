# Task Manager

This is a simple Task Manager application written in Go.

## Features

- Add new tasks
- List all tasks
- Mark tasks as completed
- Delete tasks

## Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/yourusername/task-manager.git
    ```
2. Navigate to the project directory:
    ```sh
    cd task-manager
    ```
3. Build the application:
    ```sh
    go build
    ```

## Usage

Run the application:
```sh
./task-manager
```

## Project Structure

```
task-manager/
├── cmd/server
│   ├── main.go
├── internal/handler/
│   ├── task_handler.go
│   └── task_handler_test.go
├── README.md
└── go.mod
```

- `main.go`: Entry point of the application.
- `task/task.go`: Contains the task management logic.
- `task/task_test.go`: Contains tests for the task management logic.
- `go.mod`: Go module file.

## Contributing

1. Fork the repository.
2. Create a new branch (`git checkout -b feature-branch`).
3. Make your changes.
4. Commit your changes (`git commit -am 'Add new feature'`).
5. Push to the branch (`git push origin feature-branch`).
6. Create a new Pull Request.

## License

This project is licensed under the MIT License.