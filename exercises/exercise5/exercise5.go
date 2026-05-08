package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	title       string
	description string
	completed   bool
}

const addTaskOption = 1
const markCompletedOption = 2
const editTaskOption = 3
const deleteTaskOption = 4
const cancel = 0
const titleChange = 1
const descriptionChange = 2

func printTasks(tasks []Task) {
	fmt.Printf("|| My Tasks ||\n")
	for _, task := range tasks {
		fmt.Printf("Title: %s\n", task.title)
		fmt.Printf("Description: %s\n", task.description)
		if task.completed {
			fmt.Printf("Completed: ✓\n")
		} else {
			fmt.Printf("Completed: ✗\n")
		}
		fmt.Printf("-------------\n")
	}
}

func NewTask(title, description string) Task {
	return Task{
		title:       title,
		description: description,
		completed:   false,
	}
}

func findTaskByTitle(tasks []Task, title string) (*Task, error) {
	for i := range tasks {
		if tasks[i].title == title {
			return &tasks[i], nil
		}
	}
	return nil, fmt.Errorf("task not found")
}

func (t *Task) MarkCompleted() {
	t.completed = true
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	tasks := []Task{}

	for {
		fmt.Println("| Task Manager |")
		printTasks(tasks)
		fmt.Println("[1] Add Task")
		fmt.Println("[2] Mark Task as Completed")
		fmt.Println("[3] Edit Task")
		fmt.Println("[4] Delete Task")
		fmt.Println("[0] Exit")
		fmt.Println("Select an option:")
		optionStr, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}
		optionStr = strings.TrimSpace(optionStr)
		option, err := strconv.Atoi(optionStr)
		if err != nil {
			fmt.Println("Invalid option. Please enter a number.")
			continue
		}

		switch option {
		case addTaskOption:
			var title, description string
			fmt.Println("Enter task title:")
			title, _ = reader.ReadString('\n')
			title = strings.TrimSpace(title)
			if title == "" {
				fmt.Println("Title cannot be empty.")
				break
			}
			if _, err := findTaskByTitle(tasks, title); err == nil {
				fmt.Println("Task with this title already exists. Please choose a different title.")
				break
			}
			fmt.Println("Enter task description:")
			description, _ = reader.ReadString('\n')
			description = strings.TrimSpace(description)
			tasks = append(tasks, NewTask(title, description))
		case markCompletedOption:
			var title string
			fmt.Println("Enter task title to mark as completed:")
			title, _ = reader.ReadString('\n')
			title = strings.TrimSpace(title)
			if title == "" {
				fmt.Println("Title cannot be empty.")
				break
			}
			task, err := findTaskByTitle(tasks, title)
			if err == nil {
				task.MarkCompleted()
				fmt.Println("Task marked as completed.")
			} else {
				fmt.Println("Task not found.")
			}
		case editTaskOption:
			var title string
			fmt.Println("Enter task title to edit:")
			title, _ = reader.ReadString('\n')
			title = strings.TrimSpace(title)
			if title == "" {
				fmt.Println("Title cannot be empty.")
				break
			}
			task, err := findTaskByTitle(tasks, title)
			if err == nil {
				fmt.Println("-------------")
				fmt.Println("Title:", task.title)
				fmt.Println("Description:", task.description)
				fmt.Println("Completed:", task.completed)
				fmt.Println("-------------")
				fmt.Println()
				fmt.Println("What do you want to change?")
				fmt.Println("[1] Title")
				fmt.Println("[2] Description")
				fmt.Println("[0] Cancel")
				changeStr, _ := reader.ReadString('\n')
				changeStr = strings.TrimSpace(changeStr)
				change, _ := strconv.Atoi(changeStr)
				switch change {
				case titleChange:
					var newTitle string
					fmt.Println("Enter new title:")
					newTitle, _ = reader.ReadString('\n')
					newTitle = strings.TrimSpace(newTitle)
					if newTitle == "" {
						fmt.Println("Title cannot be empty.")
						break
					}
					if _, err := findTaskByTitle(tasks, newTitle); err == nil && newTitle != task.title {
						fmt.Println("Task with this title already exists. Please choose a different title.")
						break
					}
					task.title = newTitle
				case descriptionChange:
					var newDescription string
					fmt.Println("Enter new description:")
					newDescription, _ = reader.ReadString('\n')
					newDescription = strings.TrimSpace(newDescription)
					task.description = newDescription
				case cancel:
					fmt.Println("Edit cancelled.")
				default:
					fmt.Println("Invalid option.")
				}
			} else {
				fmt.Println("Task not found.")
			}
		case deleteTaskOption:
			var title string
			fmt.Println("Enter task title to delete:")
			title, _ = reader.ReadString('\n')
			title = strings.TrimSpace(title)
			if title == "" {
				fmt.Println("Title cannot be empty.")
				break
			}
			for i, task := range tasks {
				if task.title == title {
					tasks = append(tasks[:i], tasks[i+1:]...)
					fmt.Println("Task deleted.")
					break
				}
			}
			fmt.Println("Task not found.")
		case cancel:
			fmt.Println("Exiting the task manager.")
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}
