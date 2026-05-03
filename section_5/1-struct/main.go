package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID        int
	FirstName string
	LastName  string
	Position  string
	Salary    int
	IsActive  bool
	JoinedAt  time.Time
}

func NewEmployee(id int, firstName, lastName, positions string, isActive bool) Employee {
	return Employee{
		ID:        id,
		FirstName: firstName,
		LastName:  lastName,
		Position:  positions,
		IsActive:  isActive,
		JoinedAt:  time.Now(),
	}
}

func main() {
	employee := Employee{
		ID:        1,
		FirstName: "John",
		LastName:  "Doe",
		Position:  "Software Engineer",
		Salary:    80000,
		IsActive:  true,
		JoinedAt:  time.Now(),
	}

	fmt.Printf("%+v\n", employee)

	jane := NewEmployee(2, "Jane", "Smith", "Product Manager", true)
	fmt.Printf("%+v\n", jane)
	jane.Salary = 9890
	fmt.Printf("%+v\n", jane)
}
