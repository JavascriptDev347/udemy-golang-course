package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {
	text1 := "Hello world! Welcome to Go"
	regGo, err := regexp.Compile(`Go`)
	if err != nil {

		fmt.Println("Error:", err)
		os.Exit(1) // stop the code
	}

	fmt.Printf("Text '%s', matches 'Go': %t\n", text1, regGo.MatchString(text1))
}
