package main

import "fmt"

// simple function without pointer
func modifyValue(val int) {
	val = val * 10
	fmt.Printf("Modified value: %d\n", val)
}

func modifyPointer(val *int) {
	if val == nil {
		fmt.Printf("Value: %d\n", val)
		return
	}
	*val = *val * 10
	fmt.Printf("Modified pointer value: %d\n", *val)

}

func main() {
	age := 10
	agePtr := &age
	fmt.Println("Age address: ", &age)
	fmt.Println("Age address: ", agePtr)

	val := 10
	modifyValue(val)
	fmt.Println("Original value: ", val)

	val2 := 12
	modifyPointer(&val2)
	fmt.Println("Original value: ", val2)
}
