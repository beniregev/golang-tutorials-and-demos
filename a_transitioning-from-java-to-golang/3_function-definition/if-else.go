package main

import "fmt"

func printParity(x int) {
	if x%2 == 0 {
		fmt.Printf("%v is even.\n", x)
	} else {
		fmt.Printf("%v is odd.\n", x)
	}
}

func printParity_no_else(x int) {
	if x%2 == 0 {
		fmt.Printf("%v is even.\n", x)
		return
	}
	fmt.Printf("%v is odd.\n", x)
}

func printParity_shorthand(x int) {
	if r := x % 2; r == 0 {
		fmt.Printf("%v is even.\n", x)
		return
	}
	fmt.Printf("%v is odd.\n", x)
}

func main() {
	printParity(4)
	printParity(5)

	printParity_no_else(6)
	printParity_no_else(7)

	printParity_shorthand(8)
	printParity_shorthand(9)
}
