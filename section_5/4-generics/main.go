package main

import "fmt"

func Sum[T int | float32 | float64](numbers ...T) T {
	var total T
	for _, n := range numbers {
		total += n
	}
	return total
}

func main() {
	fmt.Println("Generics in Go")

	i := Sum(1, 2, 3)
	fmt.Printf("First %T\n", i)
	f := Sum(1.5, 2.5, 3.5)
	fmt.Printf("Second %T\n", f)

}
