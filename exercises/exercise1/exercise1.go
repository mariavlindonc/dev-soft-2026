package main

import (
	"fmt"
	"math"
)

func main() {
	// declare variables
	var side1, side2, hypotenuse, area, perimeter float64

	// prompt user for sides of the triangle
	fmt.Print("Please, enter the length of the first side: ")
	fmt.Scanln(&side1)
	fmt.Print("Please, enter the length of the second side: ")
	fmt.Scanln(&side2)

	// calculate the hypotenuse using the Pythagorean theorem
	hypotenuse = math.Hypot(side1, side2)

	// calculate the area of the triangle
	area = (side1 * side2) / 2

	// calculate the perimeter of the triangle
	perimeter = side1 + side2 + hypotenuse

	// print the results
	fmt.Printf("The area of the triangle is: %.2f\n", area)
	fmt.Printf("The perimeter of the triangle is: %.2f\n", perimeter)
}
