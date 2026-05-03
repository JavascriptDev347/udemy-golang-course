package main

import (
	"errors"
	"fmt"
	"strings"
)

func divide(a, b float64) (float64, error) {
	if b == 0 {
		fmt.Println("divide by zero")
		return 0, errors.New("divide by zero")
	}
	return a / b, nil
}

func splitName(fullName string) (firstName, lastName string) {
	parts := strings.Split(fullName, " ")
	firstName = parts[0]
	lastName = parts[1]

	return
}

func main() {
	result, err := divide(2.0, 3.0)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)

	firsName, lastName := splitName("John Doe")
	fmt.Printf("First Name: %s, Last Name: %s\n", firsName, lastName)
}
