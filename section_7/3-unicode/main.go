package main

import (
	"fmt"
	"unicode"
)

func main() {
	username := "Russel"
	fmt.Println(len(username))
	fmt.Printf("%c\n", username[0])

	data := []rune{'안', '녕', '하', '세', '요'}
	for i, datum := range data {
		fmt.Println(string(datum))
		fmt.Printf("%d: %c\n", i, datum)
	}

	u := unicode.IsDigit(data[0])
	fmt.Println(u)
}
