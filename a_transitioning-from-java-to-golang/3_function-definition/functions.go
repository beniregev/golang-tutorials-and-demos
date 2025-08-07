package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func main() {
	x := 3
	y := 2

	sub := func(a int, b int) int {
		return a - b
	}

	fmt.Printf("add(%v,%v): %v\n", x, y, add(x, y))
	fmt.Printf("sub(%v,%v): %v\n", x, y, sub(x, y))
}
