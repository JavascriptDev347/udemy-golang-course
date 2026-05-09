package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	tempFile, err := os.CreateTemp("", "logs.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		fmt.Println("removing tempFile", tempFile.Name())
		//_ = os.Remove(tempFile.Name())
	}()

	_, err = tempFile.Write([]byte("Hello world !!!"))
	if err != nil {
		log.Fatal(err)
		tempFile.Close()
		return
	}
}
