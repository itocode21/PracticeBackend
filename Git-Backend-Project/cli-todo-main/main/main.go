package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"
)

// ---------------------------------------------
// task struct
type Task struct {
	ID        int       `json:"id"`
	Text      string    `json:"text"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_time"`
	UpdatedAt time.Time `json:"update_time"`
}

var tasks []Task

// ---------------------------------------------
// read json file
func loadTasks() {
	file, err := os.ReadFile("tasks.json")
	if err != nil {
		return
	}
	json.Unmarshal(file, &tasks)
}

// ---------------------------------------------
// write task in file
func saveTasks() {
	file, _ := json.MarshalIndent(tasks, "", "    ")
	_ = os.WriteFile("tasks.json", file, 0644)
}

// ---------------------------------------------
// add new task || add "u task" || set default status "In progress"
func addTask(text string) {
	task := Task{
		ID:   len(tasks) + 1,
		Text: text, Status: "In progress",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now()}
	tasks = append(tasks, task)
	saveTasks()
}

// ---------------------------------------------
// update task write new texr in created task || update [task index] "new"
func updateTask(id int, newText string) {
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Text = newText
			tasks[i].UpdatedAt = time.Now()
			saveTasks()
			return
		}
	}

}

// ---------------------------------------------
// delete task || delete [task index]
func deleteTask(id int) {
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			saveTasks()
			return
		}
	}
}

// ---------------------------------------------
// update task status || done\cancel\pending [task id]
func updateTaskStatus(id int, newStatus string) {
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Status = newStatus
			saveTasks()
			return
		}
	}
}

// ---------------------------------------------
// list all task || list
func listTasks(status string) {
	for _, task := range tasks {
		if status == "" || task.Status == status {
			fmt.Printf("[%d] %s --> Status:%s\n | Created at: %s\n | Updated at: %s\n",
				task.ID, task.Text, task.Status,
				task.CreatedAt.Format("[date:2006-01-02] [time:15:04:05]"),
				task.UpdatedAt.Format("[date:2006-01-02] [time:15:04:05]"))
		}
	}
}

// ---------------------------------------------
// list all task with status "done"|| list-done
func listDone() {
	for _, task := range tasks {
		if task.Status == "Done" {
			fmt.Printf("[%d] %s --> Status:%s\n | Created at: %s\n | Updated at: %s\n",
				task.ID, task.Text, task.Status,
				task.CreatedAt.Format("[date:2006-01-02] [time:15:04:05]"),
				task.UpdatedAt.Format("[date:2006-01-02] [time:15:04:05]"))
		}
	}
}

// ---------------------------------------------
// list all task with status "inprogress"|| list-progress
func listPending() {
	for _, task := range tasks {
		if task.Status == "Pending" {
			fmt.Printf("[%d] %s --> Status:%s\n | Created at: %s\n | Updated at: %s\n",
				task.ID, task.Text, task.Status,
				task.CreatedAt.Format("[date:2006-01-02] [time:15:04:05]"),
				task.UpdatedAt.Format("[date:2006-01-02] [time:15:04:05]"))
		}
	}
}

// ---------------------------------------------
// list all task with status "cancel"|| list-cancel
func listCancel() {
	for _, task := range tasks {
		if task.Status == "Cancel" {
			fmt.Printf("[%d] %s --> Status:%s\n | Created at: %s\n | Updated at: %s\n",
				task.ID, task.Text, task.Status,
				task.CreatedAt.Format("[date:2006-01-02] [time:15:04:05]"),
				task.UpdatedAt.Format("[date:2006-01-02] [time:15:04:05]"))
		}
	}
}

func main() {
	loadTasks()
	if len(os.Args) < 2 {
		fmt.Println("No command provided")
		return
	}
	command := os.Args[1]
	switch command {
	//add new task
	case "add":
		text := os.Args[2]
		addTask(text)
	//update created task
	case "update":
		id, _ := strconv.Atoi(os.Args[2])
		newText := os.Args[3]
		updateTask(id, newText)
	//delete task
	case "delete":
		id, _ := strconv.Atoi(os.Args[2])
		deleteTask(id)
	//list all task
	case "list":
		status := ""
		if len(os.Args) > 2 {
			status = os.Args[2]
		}
		listTasks(status)
	//set done status task
	case "done":
		id, _ := strconv.Atoi(os.Args[2])
		updateTaskStatus(id, "Done")
	//set pending status task
	case "pending":
		id, _ := strconv.Atoi(os.Args[2])
		updateTaskStatus(id, "Pending")
	//set cancel status task
	case "cancel":
		id, _ := strconv.Atoi(os.Args[2])
		updateTaskStatus(id, "Cancel")
		//list only one status task
	case "list-done":
		listDone()
	case "list-pending":
		listPending()
	case "list-cancel":
		listCancel()

	default:
		fmt.Println("Invalid command")
	}
}
