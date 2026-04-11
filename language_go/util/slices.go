package util

import "fmt"

func PrintSlices() {
	// Slices in Go are flexible arrays.

	// Declare a slice of integers
	daysOfTheWeek := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

	// Create a slice from another slice
	weekEnds := daysOfTheWeek[5:7]

	fmt.Println(len(weekEnds)) // Indicates the number of elements in the slice
	fmt.Println(cap(weekEnds)) // Indicates the total capacity of the slice (the number of elements in the underlying array starting from the first element of the slice)

	weekEnds = append(weekEnds, "NewDay") // Appending an element to the slice

	fmt.Println(cap(weekEnds)) // appending may increase the capacity of the slice if it exceeds the current capacity.
	// When exceeding the current capacity, Go will allocate a new underlying array and copy the existing elements to it.

	names := make([]string, 5, 10) // Creates a slice of strings with length 5 and capacity 10
	names[0] = "Alice"             // 5 elements can be initialized this way and the rest can be initialized using append

	slice1 := []int{25, 30, 35, 40, 45}
	slice2 := make([]int, len(slice1)) // Creates a slice of integers with length 5 and capacity 5
	fmt.Println(slice2)                // Output: [0 0 0 0 0] - slice2 is initialized with zero values

	copy(slice2, slice1) // Copies elements from slice1 to slice2
	fmt.Println(slice2)  // Output: [25 30 35 40 45] - slice2 now contains the same elements as slice1
}
