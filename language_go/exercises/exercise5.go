package exercises

import "fmt"

type Task struct {
	title       string
	description string
	completed   bool
}

func ManageTasks() {
	option := 0

	for {
		fmt.Println("| Task Manager |")
		fmt.Println("[1] Add Task")
		fmt.Println("[2] Mark Task as Completed")
		fmt.Println("[3] Edit Task")
		fmt.Println("[4] Delete Task")
		fmt.Println("[0] Exit")
		fmt.Println("Select an option:")
		fmt.Scanln(&option)

		switch option {
		case 1:
			// Code to add a task
		case 2:
			// Code to mark a task as completed
		case 3:
			// Code to edit a task
		case 4:
			// Code to delete a task
		case 0:
			fmt.Println("Exiting the task manager.")
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}
