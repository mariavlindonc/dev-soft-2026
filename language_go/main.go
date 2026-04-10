package main

import (
	"fmt"
	"language_go/exercise1"
	"language_go/exercise4"
)

func main() {
	option := 0

	fmt.Println("Welcome to Go programming!")
	fmt.Println("[1] Exercise 1")
	fmt.Println("[2] Exercise 2")
	fmt.Println("[3] Exercise 3")
	fmt.Println("[4] Exercise 4")
	fmt.Println("[0] Exit")
	fmt.Scanln(&option)

	switch option {
	case 1:
		fmt.Println("Starting Exercise 1...")
		exercise1.AreaAndPerimeter()
	case 2:
		fmt.Println("Starting Exercise 2...")
		// exercise2()
	case 3:
		fmt.Println("Starting Exercise 3...")
		// exercise3()
	case 4:
		fmt.Println("Starting Exercise 4...")
		exercise4.DisplayMenu()
	case 0:
		fmt.Println("Exiting the program.")
	default:
		fmt.Println("Invalid option. Please try again.")
	}

	if option != 0 {
		main() // Restart the menu after completing an exercise
	}
}
