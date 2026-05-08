package util

import "fmt"

func PrintIteration() {
	// Create a slice of integers
	numbers := []int{1, 2, 3, 4, 5}

	// Classic for loop
	// Used when you need to use an index to access elements or when you need to modify the loop variable.
	for i := 0; i < len(numbers); i++ {
		fmt.Printf("Index: %d, Value: %d\n", i, numbers[i])
	}

	// Range-based for loop
	// Used when you want to iterate over all elements in a slice or map without needing the index.
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}
}
