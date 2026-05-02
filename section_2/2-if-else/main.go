package main

import "fmt"

func main() {
	tmp := 25
	if tmp > 30 {
		fmt.Println("HOT")
	} else {
		fmt.Println("NORM")
	}

	score := 87
	if score >= 90 {
		fmt.Println("A")
	} else if score >= 80 {
		fmt.Println("B")
	} else if score >= 70 {
		fmt.Println("C")
	} else {
		fmt.Println("Failed")
	}

	userAccess := map[string]bool{
		"Russel": true,
		"James":  true,
		"Goose":  false,
	}

	if hasAccess, ok := userAccess["Russel"]; ok {
		if hasAccess {
			fmt.Println(hasAccess, ok)
		}
	}
}
