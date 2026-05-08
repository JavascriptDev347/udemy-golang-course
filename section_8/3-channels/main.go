package main

import (
	"fmt"
	"time"
)

type user struct {
	Name string
}

func main() {
	messages := make(chan string)
	users := make(chan user)

	go func() {
		fmt.Println("Sending message...")
		messages <- "ping"
	}()

	go func() {
		fmt.Println("Sending message 2 ...")
		messages <- "ping 2"
	}()

	go func() {
		fmt.Println("Sending message 3 ...")
		users <- user{
			Name: "Go Routine 1",
		}
	}()

	time.Sleep(1 * time.Second)
	msg := <-messages
	fmt.Println(msg)

	msg = <-messages
	fmt.Println(msg)

	user := <-users
	fmt.Println(user)

}
