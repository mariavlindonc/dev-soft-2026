package util

import "fmt"

func PrintDataTypes() {
	var age int = 30                 // Integer data type
	var visits int64 = 10000000000   // 64-bit integer data type
	var temperature32 float32 = 36.5 // 32-bit floating-point data type
	var temperature64 float64 = 36.5 // 64-bit floating-point data type
	var name string = "John Doe"     // String data type
	var isActive bool = true         // Boolean data type

	fmt.Printf("Age: %d\n", age)
	fmt.Printf("Visits: %d\n", visits)
	fmt.Printf("Temperature (32-bit): %.2f\n", temperature32)
	fmt.Printf("Temperature (64-bit): %.2f\n", temperature64)
	fmt.Printf("Name: %s\n", name)
	fmt.Printf("Is Active: %t\n", isActive)
}
