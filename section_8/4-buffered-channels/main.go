package main

import "fmt"

func main() {
	messages := make(chan string, 3)

	fmt.Println("Sending messages to buffered channel")
	messages <- "First Message"
	messages <- "Second Message"
	messages <- "Third Message"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
