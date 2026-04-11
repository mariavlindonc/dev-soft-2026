package util

import "fmt"

func PrintLogicalOperators() {
	// Logical operators in Go are used to combine multiple boolean expressions.
	// The main logical operators are: && (logical AND), || (logical OR), and ! (logical NOT).

	age := 25
	isAdult := true
	hasLicense := false
	accountBalance := 1500.0

	// Using logical AND (&&) to check if a person can drive
	if age >= 18 && hasLicense {
		fmt.Println("Can drive.")
	} else {
		fmt.Println("Cannot drive.")
	}

	// Using logical OR (||) to check if a person is eligible for a credit card
	if isAdult || accountBalance > 1000 {
		fmt.Println("Eligible for credit card.")
	} else {
		fmt.Println("Not eligible for credit card.")
	}

	// Using logical NOT (!) to check if a person does not have a driver's license
	if !hasLicense {
		fmt.Println("Needs to obtain a driver's license.")
	}
}
