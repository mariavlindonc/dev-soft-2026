package util

import "fmt"

func PrintLogicalOperators() {
	age := 25
	isAdult := true
	hasLicense := false
	accountBalance := 1500.0

	if age >= 18 && hasLicense {
		fmt.Println("Can drive.")
	} else {
		fmt.Println("Cannot drive.")
	}

	if isAdult || accountBalance > 1000 {
		fmt.Println("Eligible for credit card.")
	} else {
		fmt.Println("Not eligible for credit card.")
	}

	if !hasLicense {
		fmt.Println("Needs to obtain a driver's license.")
	}
}
