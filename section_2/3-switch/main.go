package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to Switch Program")

	today := "Sunday"

	switch today {
	case "Sunday":
		fmt.Println("Sunday day offf")

	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("Working day again :cry")
	default:
		fmt.Println("Saturday")
	}

	switch hour := time.Now().Hour(); {
	case hour < 12:
		fmt.Println("Good morning")
	case hour < 17:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good evening")
	}

	checkType := func(i interface{}) {
		switch v := i.(type) {
		case int:
			fmt.Printf("Twice %v is %v\n", v, v*2)
		case string:
			fmt.Printf("%q is a string\n", v)
		case bool:
			fmt.Printf("Twice %v is a bool\n", v)
		default:
			fmt.Printf("I don't know about type %T!\n", v)
		}
	}

	checkType(2)
	checkType("Hello")

}
