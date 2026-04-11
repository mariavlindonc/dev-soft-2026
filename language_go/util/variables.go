package util

import "fmt"

func PrintVariables() {
	// For variable declaration, you can use the var keyword followed by the variable name and type.
	var name string
	name = "Alice"
	fmt.Println("Name:", name)
	// You can also update the value of a variable after it has been declared.
	name = "Bob"
	fmt.Println("Updated Name:", name)

	// You can also declare and initialize a variable in one line.
	var city string = "New York"
	fmt.Println("City:", city)

	// For short variable declaration, you can use the := syntax, which infers the type from the assigned value.
	age := 30
	fmt.Println("Age:", age)
}
