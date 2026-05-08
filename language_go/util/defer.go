package util

import "fmt"

func PrintDefer() {
	// Defer statements are executed in LIFO order (Last In, First Out)
	// They are used to postpone the execution of a function until the surrounding function returns.
	// It is used to:
	// - Administrate resources (e.g., closing files, releasing locks)
	// - Ensure that certain operations are performed (e.g., logging, cleanup)
	// - Handle errors (e.g., recovering from panics)
	defer fmt.Println("3")
	fmt.Println("1")
	fmt.Printf("2")
}
