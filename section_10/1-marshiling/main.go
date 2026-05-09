package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	Name     string `json:"name" xml:"name"`
	Email    string `json:"email" xml:"email"`
	Age      int    `json:"age" xml:"age"`
	IsActive bool   `json:"isActive" xml:"isActive"`
}

func main() {
	fmt.Println("Marshiling...")

	user := user{
		Name:     "Russel",
		Email:    "russel@gmail.com",
		Age:      20,
		IsActive: true,
	}

	byteSlices, err := json.MarshalIndent(user, " ", " ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(byteSlices))
}
