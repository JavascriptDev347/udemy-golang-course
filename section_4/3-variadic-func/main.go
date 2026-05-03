package main

import "fmt"

func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func config(numbers ...int) {
	if len(numbers) > 0 {
		first := numbers[0]
		fmt.Println("First number:", first)
	} else {
		fmt.Println("Default value")
	}
}

func main() {
	result := sum(1, 2, 3, 4)
	fmt.Println(result)

	config()
}
