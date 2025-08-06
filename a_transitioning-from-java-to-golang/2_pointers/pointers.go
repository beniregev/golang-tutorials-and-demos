package main

import "fmt"

func main() {
	// Declare pointer
	var ptr *string

	// initialize a greeting
	greeting := "Hello, world!"

	// Assign greeting address to pointer
	ptr = &greeting

	// Print out our variables
	fmt.Println("Greeting: ", greeting)
	fmt.Println("Address of greeting: ", ptr)
	fmt.Println("Value stored in ptr: ", *ptr)
}
