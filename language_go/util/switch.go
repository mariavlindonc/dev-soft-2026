package util

import "fmt"

func PrintSwitchExample() {
	day := "Monday"

	switch day {
	case "Monday":
		fmt.Println("Start of the week.")
	case "Friday":
		fmt.Println("End of the week.")
	default:
		fmt.Println("Midweek.")
	}

	//Using switch is ideal when you need to evaluate an expression against multiple predefined values.
	// It is often preferred over a series of if-else statements for better readability and maintainability.
	// It's ideal for single variable comparisons, while if-else statements are better for complex conditions or multiple variables.
}
