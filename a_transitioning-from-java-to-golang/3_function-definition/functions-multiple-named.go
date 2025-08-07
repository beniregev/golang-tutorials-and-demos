package main

import "fmt"

func rectangle(x int, y int) (area int, circumf int) {
	area = x * y
	circumf = (x + y) * 2
	return
}

func main() {
	a := 2
	b := 3
	area, circumf := rectangle(a, b)
	fmt.Printf("rectangle(%v, %v): area = %v;\n", a, b, area)
	fmt.Printf("rectangle(%v, %v): circumference = %v;\n", a, b, circumf)
	fmt.Printf("rectangle(%v, %v): area = %v; circumference = %v;\n", a, b, area, circumf)
}
