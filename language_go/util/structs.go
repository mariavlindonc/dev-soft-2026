package util

import "fmt"

type Person struct {
	name string
	age  int
}

func PrintStructs() {
	// Create an instance of the Person struct
	var person1 Person
	// Assign values to the fields of the struct
	person1.name = "Alice"
	person1.age = 30

	// Create an instance of the Person struct using a struct literal
	person2 := Person{name: "Bob", age: 25}

	// Print the details of the persons
	fmt.Println(person1.name)
	fmt.Println(person1.age)
	fmt.Printf("Name: %s, Age: %d\n", person2.name, person2.age)
}
