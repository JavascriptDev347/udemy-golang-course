package main

import (
	"fmt"
	"time"
)

func sayHello(text string, delay time.Duration) {
	time.Sleep(delay)
	fmt.Println("Say Hello: ", text)
}

func main() {
	fmt.Println("Hello from main goroutine")

	go sayHello("Go Routine 1", time.Second)

	fmt.Println("Last message from main goroutine")
	time.Sleep(2 * time.Second)
}
