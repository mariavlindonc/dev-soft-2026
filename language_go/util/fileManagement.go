package util

import (
	"fmt"
	"os"
)

const path = "data.txt"

func CreateFile() {
	// Create a file named "data.txt" in the current directory
	file, _ := os.Create(path) // Discard the error for simplicity
	file.Close()               // Make sure it's closed after creation
	fmt.Println("File created successfully.")
}

func WriteFile() {
	file, _ := os.Create(path)

	// Defer the closing
	defer file.Close()

	// Write some content to the file
	file.WriteString("Hello, this is a sample text written to the file.\n")
}

func ReadFile() {
	CreateFile()             // Ensure the file exists before reading
	file, _ := os.Open(path) // Open the file for reading
	defer file.Close()       // Defer the closing

	// Write something to the file before reading
	file.Write([]byte("Hello, sample text for reading.\n")) // []byte is used to write string data

	// Read the content of the file
	data, _ := os.ReadFile(path)

	// Print the variable in which the content of the file is stored
	fmt.Print("Content of the file:", string(data))
}
