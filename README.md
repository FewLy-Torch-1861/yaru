# Yaru (やることリスト)

> [!WARNING]
> **Note**: This project is currently a **Work In Progress (WIP)**. Features and commands may change as development continues.

**Yaru** is a minimalist TUI (Text User Interface) to-do list manager written in Go. The name comes from "Yarukotorisuto" (やることリスト), meaning "to-do list" in Japanese.

It is designed to be a simple, lightweight tool for managing tasks directly from your terminal.

## Features

- **Add Tasks**: Quickly add new tasks with a description.
- **List Tasks**: View all your tasks, including their status (✅/❌) and creation date.
- **Mark as Done**: Complete tasks by their ID.
- **Persistent Storage**: Tasks are saved to a JSON file (currently located in `tmp/yaru/tasks.json`).
- **Minimalist Design**: Clean output with basic color coding.

## Prerequisites

- **Go**: You need Go installed on your machine (specifically Go 1.25.5 or later, as per `go.mod`).

## Installation & Setup

1.  **Clone the repository:**
    ```bash
    git clone https://github.com/FewLy-Torch-1861/yaru.git
    cd yaru
    ```

2.  **Build the application:**
    ```bash
    go build -o yaru main.go
    ```

3.  **Run directly (optional):**
    You can run it without building using `go run main.go`.

## Usage

### Add a Task
Add a new task by providing a description.

```bash
./yaru add "Buy groceries"
```

### List Tasks
Display all tasks in a formatted list.

```bash
./yaru list
```
*Output Example:*
```text
#1  [❌] Buy groceries       (23-12-29)
#2  [✅] Walk the cat        (23-12-28)
```

### Mark a Task as Done
Mark a task as completed using its ID (found via the `list` command).

```bash
./yaru done 1
```

## Contributing

Contributions are welcome! Since this is a learning project, feel free to open issues or submit pull requests for:
- Bug fixes
- Code refactoring
- New feature suggestions

1.  Fork the repository.
2.  Create your feature branch (`git checkout -b feature/amazing-feature`).
3.  Commit your changes (`git commit -m 'Add some amazing feature'`).
4.  Push to the branch (`git push origin feature/amazing-feature`).
5.  Open a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
