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

// FullName value receiver funct => this function can't edit main employye only for read
func (e *Employee) FullName() string {
	if e == nil {
		return "Unknown Employee"
	}
	return e.FirstName + " " + e.LastName
}

func (e *Employee) EditActive(active bool) bool {
	e.IsActive = active
	return e.IsActive
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
	employees := []Employee{}
	employee := NewEmployee(2, "Jane", "Smith", "Product Manager", true)
	employee2 := NewEmployee(3, "Jane 2", "Smith 2", "Team Leader", false)
	employees = append(employees, employee)
	employees = append(employees, employee2)

	fmt.Println(employee.FullName())
	fmt.Println("all employees jane before edit", employees)
	fmt.Printf("Before Edit: %v\n", employee.IsActive)
	employee.EditActive(false)
	fmt.Printf("After Edit: %v\n", employee.IsActive)

	fmt.Println("all employees", employees)

}
