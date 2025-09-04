# 📝 Task Tracker CLI

A simple command-line interface (CLI) tool built in **Go** to track and manage your tasks.  
It helps you keep track of what you need to do, what you’re working on, and what you’ve completed.  

All tasks are stored in a `tasks.json` file in the project directory, so they persist between runs.  

https://roadmap.sh/projects/task-tracker

---

## 🚀 Features
- Add new tasks with description and status
- Update existing tasks (description/status)
- Delete tasks
- Mark tasks as `in-progress` or `done`
- List all tasks
- List tasks filtered by status (`todo`, `in-progress`, `done`)
- JSON-based persistence (no external DB required)

---

## 📦 Installation

1. Clone this repository:
   ```bash
     git clone https://github.com/A-Aljami/Task-Tracker.git
   cd task-cli
   ```

2. Run the program with Go:
   ```bash
   go run main.go
   ```

Or build a binary for faster use:
   ```bash
   go build -o task-cli
   ./task-cli
   ```

---

## ⚡ Usage

Here are all the available commands:

### Add a task
```bash
task-cli add "Buy groceries"
task-cli add "Finish Go project" in-progress
```
✅ Creates a task with description and optional status (`todo` by default).

---

### List tasks
```bash
task-cli list
task-cli list todo
task-cli list in-progress
task-cli list done
```
📋 Shows all tasks or filters them by status.

---

### Update a task
```bash
task-cli update 1 "Buy groceries and cook dinner"
task-cli update 1 "Complete Go CLI project" done
```
✏️ Updates a task’s description and/or status.

---

### Delete a task
```bash
task-cli delete 1
```
🗑 Deletes the task with the given ID.

---

### Mark tasks
```bash
task-cli mark-in-progress 2
task-cli mark-done 3
```
⚡ Quickly updates a task’s status without changing the description.

---

## 📂 Task Structure

Each task is stored in `tasks.json` with the following properties:
```json
{
  "id": 1,
  "description": "Buy groceries",
  "status": "todo",
  "createdAt": "2025-09-04T17:10:00Z",
  "updatedAt": "2025-09-04T17:10:00Z"
}
```

---

## ✅ Example Workflow

```bash
# Add tasks
task-cli add "Learn Go"
task-cli add "Build CLI app" in-progress

# List tasks
task-cli list
task-cli list in-progress

# Update a task
task-cli update 1 "Learn Go deeply" done

# Mark quickly
task-cli mark-done 2

# Delete a task
task-cli delete 1
```

---

## 🛠 Tech Stack
- [Go](https://golang.org/) (standard library only)
- JSON file for persistence (`tasks.json`)

---

## 📌 Notes
- If `tasks.json` does not exist, it will be created automatically.
- Only **native Go packages** are used (no external dependencies).
