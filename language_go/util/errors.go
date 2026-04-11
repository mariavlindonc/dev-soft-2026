package util

import (
	"errors"
	"fmt"
)

// Errors are values that indicate that something went wrong during the execution of a program.
// They are used to handle exceptional situations and provide feedback to the user or developer about what went wrong.
// In Go, try-catch is not used for error handling.
// error is a built-in data type in Go, implemented through the errors package.

// Panic stops the execution of the program.
// It is used for unrecoverable errors, such as when the program encounters a critical issue that cannot be handled gracefully.
// It unpacks the call stack and prints the panic message along with the stack trace, which can be useful for debugging.
// It also executes any deferred functions before terminating the program.

// When to use each:
// - Use error for expected errors that can be handled.
// - Use panic for unrecoverable errors that should not occur.

// dividing receives two numbers and returns a result or an error
func Divide(a, b float64) (float64, error) {
	errorMessage := "Error: division by zero is not allowed."

	// Error handling happens before the execution of the main logic of the function.
	// Check if the divisor is zero.
	if b == 0 {
		return 0, errors.New(errorMessage) // errors.New creates a new error with the provided message
	}

	return a / b, nil // nil indicates that there is no error
}

func PrintErrors() {
	result, err := Divide(10, 0) // Calls the function and saves the result and error in variables.

	if err != nil { // Checks if there is an error.
		fmt.Println(err) // Prints the error message if there is an error.
		return           // Exits the function if there is an error.
	}

	fmt.Println("Result:", result) // Prints the result only if there is no error.
}
