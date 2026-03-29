package util

import "fmt"

const Pi float32 = 3.14 // Constants must be initialized at the time of declaration
const Euler = 2.71828   // It can be declared without specifying the type, and it will be inferred from the value

func PrintConstants() {
	fmt.Printf("Pi: %f\n", Pi)
	fmt.Printf("Euler's Number: %f\n", Euler)
}
