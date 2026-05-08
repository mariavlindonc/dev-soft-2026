package util

import "fmt"

func PrintArrays() {
	// Arrays in Go are fixed-size sequences of elements of the same type.
	// They are defined with a specific size and cannot be resized after declaration.
	// The elements of an array are accessed using indices, starting from 0 to n-1, where n is the size of the array.

	// Declare an array of integers with a size of 5
	var myArray [5]int

	// Initialize the array with values
	myArray[0] = 10
	myArray[1] = 20
	myArray[2] = 30
	myArray[3] = 40
	myArray[4] = 50

	// Print the array
	fmt.Println("Array:", myArray)
	fmt.Println("Length of the array:", len(myArray))
	fmt.Println("Third element:", myArray[2])
}
