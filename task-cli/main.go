package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

var tasks []Task
var nextID = 1
const tasksFile = "tasks.json"

// ------------------- File persistence -------------------

// Load tasks from JSON file
func loadTasks() {
	data, err := ioutil.ReadFile(tasksFile)
	if err != nil {
		return // file may not exist yet
	}
	json.Unmarshal(data, &tasks)

	// set nextID to continue properly
	maxID := 0
	for _, t := range tasks {
		if t.ID > maxID {
			maxID = t.ID
		}
	}
	nextID = maxID + 1
}

// Save tasks to JSON file
func saveTasks() {
	data, _ := json.MarshalIndent(tasks, "", "  ")
	ioutil.WriteFile(tasksFile, data, 0644)
}

// ------------------- CRUD functions -------------------

func addTask(desc string, status string) {
	now := time.Now()
	task := Task{
		ID:          nextID,
		Description: desc,
		Status:      status,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	tasks = append(tasks, task)
	nextID++
	saveTasks()
	fmt.Println("✅ Task added successfully (ID:", task.ID, ")")
}

func listTasks(filter string) {
	if len(tasks) == 0 {
		fmt.Println("📭 No tasks found.")
		return
	}

	fmt.Println("\n📋 Task List:")
	for _, t := range tasks {
		if filter == "" || t.Status == filter {
			fmt.Printf("[%d] %s | Status: %s | Created: %s | Updated: %s\n",
				t.ID, t.Description, t.Status,
				t.CreatedAt.Format("2006-01-02 15:04"),
				t.UpdatedAt.Format("2006-01-02 15:04"))
		}
	}
}

func updateTask(id int, newDesc string, newStatus string) {
	for i, t := range tasks {
		if t.ID == id {
			if newDesc != "" {
				tasks[i].Description = newDesc
			}
			if newStatus != "" {
				tasks[i].Status = newStatus
			}
			tasks[i].UpdatedAt = time.Now()
			saveTasks()
			fmt.Println("✅ Task updated:", id)
			return
		}
	}
	fmt.Println("❌ Task not found:", id)
}

func deleteTask(id int) {
	for i, t := range tasks {
		if t.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			saveTasks()
			fmt.Println("🗑 Task deleted:", id)
			return
		}
	}
	fmt.Println("❌ Task not found:", id)
}

// ------------------- Marking shortcuts -------------------

func markTask(id int, status string) {
	for i, t := range tasks {
		if t.ID == id {
			tasks[i].Status = status
			tasks[i].UpdatedAt = time.Now()
			saveTasks()
			fmt.Printf("✅ Task %d marked as %s\n", id, status)
			return
		}
	}
	fmt.Println("❌ Task not found:", id)
}

// ------------------- CLI entrypoint -------------------

func main() {
	loadTasks() // load tasks from file at start

	if len(os.Args) < 2 {
		fmt.Println("Usage:")
		fmt.Println("  add \"description\" [status]")
		fmt.Println("  list [status]")
		fmt.Println("  update <id> \"new description\" [status]")
		fmt.Println("  delete <id>")
		fmt.Println("  mark-in-progress <id>")
		fmt.Println("  mark-done <id>")
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("❌ Missing description")
			return
		}
		desc := os.Args[2]
		status := "todo"
		if len(os.Args) > 3 {
			status = os.Args[3]
		}
		addTask(desc, status)

	case "list":
		filter := ""
		if len(os.Args) > 2 {
			filter = os.Args[2] // todo, in-progress, done
		}
		listTasks(filter)

	case "update":
		if len(os.Args) < 4 {
			fmt.Println("❌ Usage: update <id> \"new description\" [status]")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("❌ Invalid ID")
			return
		}
		newDesc := os.Args[3]
		newStatus := ""
		if len(os.Args) > 4 {
			newStatus = os.Args[4]
		}
		updateTask(id, newDesc, newStatus)

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("❌ Usage: delete <id>")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("❌ Invalid ID")
			return
		}
		deleteTask(id)

	case "mark-in-progress":
		if len(os.Args) < 3 {
			fmt.Println("❌ Usage: mark-in-progress <id>")
			return
		}
		id, _ := strconv.Atoi(os.Args[2])
		markTask(id, "in-progress")

	case "mark-done":
		if len(os.Args) < 3 {
			fmt.Println("❌ Usage: mark-done <id>")
			return
		}
		id, _ := strconv.Atoi(os.Args[2])
		markTask(id, "done")

	default:
		fmt.Println("❌ Unknown command:", command)
	}
}
