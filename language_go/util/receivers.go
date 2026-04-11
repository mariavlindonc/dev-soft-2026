package util

// We use the Person struct from the structs.go file in this package.

// Define a method on the Person struct (receiver).
func (p Person) IsAdult() bool {
	return p.age >= 18
}

func PrintReceivers() {
	// Receivers are used to call methods on a struct instance.
	person := Person{name: "Alice", age: 30}

	isAdult := person.IsAdult()
	println(person.name, "is an adult:", isAdult)
}
