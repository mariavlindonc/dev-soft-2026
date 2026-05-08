package util

import "fmt"

// Function without return value or parameters
func greet() {
	fmt.Println("Hello, World!")
}

// Function with parameters
func add(a int, b int) {
	fmt.Printf("The sum is %d\n", a+b)
}

// Function with return value
func multiply(a int, b int) int {
	return a * b
}

// Function with multiple return values
func divide(a int, b int) (int, int) {
	return a / b, a % b
}

// Function that calls other functions
func CallFunctions() {
	greet()
	add(5, 3)
	result := multiply(4, 7)
	fmt.Printf("The product is %d\n", result)
	quotient, remainder := divide(10, 3)
	fmt.Printf("The quotient is %d and the remainder is %d\n", quotient, remainder)
}
