package main

import "fmt"

func rectangle(x int, y int) (*int, error) {
	if x == 0 || y == 0 {
		return nil, fmt.Errorf("zero area: [%v, %v]", x, y)
	}
	area := x * y
	return &area, nil
}

func main() {
	a := 2
	b := 3

	area, err := rectangle(a, b)
	if err != nil {
		fmt.Println("ERROR: ", err)
		return
	}
	fmt.Printf("rectangle(%v, %v): area = %v;\n", a, b, *area)

	area, err = rectangle(a, 0)
	if err != nil {
		fmt.Println("ERROR: ", err)
		return
	}
	fmt.Printf("rectangle(%v, %v): area = %v;\n", a, b, *area)
}
