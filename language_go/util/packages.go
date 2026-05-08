package util

// Importing multiple packages in Go is done using the import statement.
// You can import standard library packages, third-party packages, or your own packages.
import (
	"fmt"
	//"math"
	//"time"
	//etc
)

func PrintPackages() {
	fmt.Println("fmt package is used for formatted I/O operations.")
	fmt.Println("math package provides basic constants and mathematical functions.")
	fmt.Println("time package provides functionality for measuring and displaying time.")
	fmt.Println("os package provides a platform-independent interface to operating system functionality.")
	fmt.Println("io package provides basic interfaces to I/O primitives.")
	fmt.Println("json package provides functions for encoding and decoding JSON data.")
	fmt.Println("net package provides a portable interface for network I/O, including TCP/IP, UDP, domain name resolution, and Unix domain sockets.")
	fmt.Println("sort package provides primitives for sorting slices and user-defined collections.")
	fmt.Println("rand package provides functions for generating pseudo-random numbers.")
}
