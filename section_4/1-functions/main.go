package main

import (
	"fmt"
)

func greet(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

func add(x int, y int) {
	fmt.Printf("%d + %d = %d\n", x, y, x+y)
}

func calculateArea(w, h float64) float64 {
	if w < 0 || h < 0 {
		fmt.Println("Width and height must be non-negative.")
		return 0
	}
	return w * h
}

func main() {
	greet("Bob")
	add(9, 8)

	c := calculateArea(5, 3)
	fmt.Println(c)
}
