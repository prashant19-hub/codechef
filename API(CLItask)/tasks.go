package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

const tasksFile = "tasks.json"

type Task struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
}

func loadTasks() ([]Task, error) {
	var tasks []Task

	data, err := os.ReadFile(tasksFile)
	if err != nil {
		if os.IsNotExist(err) {
			return tasks, nil
		}
		return nil, err
	}

	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func saveTasks(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(tasksFile, data, 0644)
}

func addTask(title string) {
	tasks, err := loadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	newID := 1
	if len(tasks) > 0 {
		newID = tasks[len(tasks)-1].ID + 1
	}

	task := Task{
		ID:        newID,
		Title:     title,
		Completed: false,
		CreatedAt: time.Now(),
	}

	tasks = append(tasks, task)

	err = saveTasks(tasks)
	if err != nil {
		fmt.Println("Error saving task:", err)
		return
	}

	fmt.Printf("Task added with ID %d\n", newID)
}

func listTasks() {
	tasks, err := loadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}

	fmt.Println("ID\tTitle\t\t\tStatus")
	fmt.Println("--\t-----\t\t\t------")
	for _, t := range tasks {
		status := "Pending"
		if t.Completed {
			status = "Done"
		}
		fmt.Printf("%d\t%-20s\t%s\n", t.ID, t.Title, status)
	}
}

func completeTask(id int) {
	tasks, err := loadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	found := false
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Completed = true
			found = true
			break
		}
	}

	if !found {
		fmt.Println("Task with ID", id, "not found.")
		return
	}

	err = saveTasks(tasks)
	if err != nil {
		fmt.Println("Error saving tasks:", err)
		return
	}

	fmt.Println("Task", id, "marked as complete.")
}

func deleteTask(id int) {
	tasks, err := loadTasks()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}

	found := false
	newTasks := []Task{}

	for _, t := range tasks {
		if t.ID == id {
			found = true
			continue
		}
		newTasks = append(newTasks, t)
	}

	if !found {
		fmt.Println("Task with ID", id, "not found.")
		return
	}

	err = saveTasks(newTasks)
	if err != nil {
		fmt.Println("Error saving tasks:", err)
		return
	}

	fmt.Println("Task", id, "deleted.")
}
