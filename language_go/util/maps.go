package util

import "fmt"

func PrintMaps() {
	// Maps in Go are unordered collections of key-value pairs.
	colors := map[string]string{
		//"key": "value",
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
	}

	// Print the entire map
	fmt.Println(colors)

	// Print the value associated with the key "blue"
	fmt.Println(colors["blue"])

	// Add another key-value pair to the map
	colors["yellow"] = "#ffff00"

	// Delete an element from the map by key
	delete(colors, "green")
}
