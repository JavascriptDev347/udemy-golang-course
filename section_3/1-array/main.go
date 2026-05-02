package main

import "fmt"

func main() {
	var numbers [2]int

	fmt.Printf("Array elements :%v\n Array length: %v\n", numbers, len(numbers))

	primes := [4]int{2, 3, 5, 7}
	fmt.Printf("Array elements :%v\n Array length: %v\n", primes, len(primes))

	for i := 0; i < len(primes); i++ {
		fmt.Printf("Element at index %v is %v\n", i, primes[i])
	}

}
