package util

import "fmt" // fmt is used for printing and scanning input

func ReadAndPrint() {
	// declare variables
	var name string
	var price float64

	// prompt user for input and read it

	fmt.Print("Please, enter your name: ") // prompt user for name
	fmt.Scanln(&name)                      // reads the full line until the enter and saves it in the variable 'name'
	// the '&' is used to pass the address of the variable in memory

	fmt.Print("Please, enter the price: ") // prompt user for price
	fmt.Scanln(&price)                     // reads the full line until the enter and saves it in the variable 'price'

	fmt.Println("Hello, ", name, "!")                      // print the name
	fmt.Printf("The price you entered is: $%.2f\n", price) // print the price with formatting to 2 decimal places
}
