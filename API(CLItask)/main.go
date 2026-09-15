package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Usage: task-manager add \"Your task\"")
			return
		}
		taskTitle := os.Args[2]
		addTask(taskTitle)

	case "list":
		listTasks()

	case "complete":
		if len(os.Args) < 3 {
			fmt.Println("Usage: task-manager complete <id>")
			return
		}
		idStr := os.Args[2]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Println("Invalid ID. Please enter a number.")
			return
		}
		completeTask(id)

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Usage: task-manager delete <id>")
			return
		}
		idStr := os.Args[2]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			fmt.Println("Invalid ID. Please enter a number.")
			return
		}
		deleteTask(id)

	default:
		fmt.Println("Unknown command:", command)
		printUsage()
	}
}

func printUsage() {
	fmt.Println("CLI Task Manager")
	fmt.Println("Usage:")
	fmt.Println("  task-manager add \"Task title\"   - Add a new task")
	fmt.Println("  task-manager list                 - List all tasks")
	fmt.Println("  task-manager complete <id>        - Mark task as complete")
	fmt.Println("  task-manager delete <id>          - Delete a task")
}
