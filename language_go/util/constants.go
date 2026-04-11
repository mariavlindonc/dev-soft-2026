package util

import "fmt"

// Constants must be initialized at the time of declaration
const Pi float32 = 3.14

// Constants can be declared without specifying the type, and it will be inferred from the value
const Euler = 2.71828

func PrintConstants() {
	fmt.Printf("Pi: %f\n", Pi)
	fmt.Printf("Euler's Number: %f\n", Euler)
}
