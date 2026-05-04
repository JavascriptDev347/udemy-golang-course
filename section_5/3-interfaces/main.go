package main

import (
	"fmt"
	"time"
)

type Person interface {
	GetName() string
	GetID() int
}

type Employee struct {
	ID        int
	FirstName string
	LastName  string
	Position  string
	Salary    int
	IsActive  bool
	JoinedAt  time.Time
}

type BusinessEmployee struct {
	ID        int
	FirstName string
	LastName  string
	Position  string
	Salary    int
	IsActive  bool
}

func (e Employee) GetID() int {
	//TODO implement me
	panic("implement me")
}

func (b BusinessEmployee) GetID() int {
	//TODO implement me
	panic("implement me")
}

func (b BusinessEmployee) GetName() string {
	return b.FirstName + " " + b.LastName
}

func (e Employee) GetName() string {
	return e.FirstName + " " + e.LastName
}

func displayName(p Person) {
	fmt.Println(p.GetName())
}

func main() {
	joe := Employee{
		ID:        1,
		FirstName: "Joe",
		LastName:  "Smith",
		Position:  "Vibe coder",
		Salary:    9000,
		IsActive:  true,
		JoinedAt:  time.Now(),
	}

	displayName(joe)

	b := BusinessEmployee{
		ID:        1,
		FirstName: "Bob",
		LastName:  "Smith",
		Position:  "Vibe coder",
		Salary:    9000,
		IsActive:  true,
	}

	displayName(b)
}
