package util

import "fmt"

func PrintIfElse() {
	// Declare a variable
	number := 10

	// Use if-else statement to check the value of the variable
	if number > 0 {
		fmt.Println("The number is positive.")
	} else if number < 0 {
		fmt.Println("The number is negative.")
	} else {
		fmt.Println("The number is zero.")
	}
}
