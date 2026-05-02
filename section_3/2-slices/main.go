package main

import "fmt"

func main() {
	names := []string{
		"John Smith",
		"Jane Smith",
		"John Smith",
	}

	fmt.Println(names)

	items := make([]int, 3, 10)
	sliceWithoutMake := []int{}
	fmt.Println(sliceWithoutMake)
	fmt.Println(items)
}
