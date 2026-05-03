package main

import "fmt"

func main() {

	// if map create without make it will be nill first value
	studentGrades := map[string]int{
		"Jane":   62,
		"John":   86,
		"Alice":  89,
		"Russel": 98,
	}
	fmt.Println("Original map", studentGrades)

	// update
	studentGrades["Alice"] = 90
	// add
	studentGrades["Bob"] = 75

	fmt.Println("New map", studentGrades)
	// check value exist or no
	alice, ok := studentGrades["Alice"]
	if ok {
		fmt.Println("Topildi: ", alice)
	}
	if russe, ok := studentGrades["Russel"]; ok {
		fmt.Println("Russel: ", russe)
	}

	// delete
	delete(studentGrades, "John")
	fmt.Println("New map after delete john", studentGrades)

	configs := make(map[string]string)

	fmt.Println(configs)

}
