package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// Encoding is the process of putting a sequence of characters into a specialized format for efficient transmission or storage.
// It's NOT security: Encoding is not for hiding secrets (that's Encryption).
// It's for compatibility: It ensures that different systems can correctly interpret and process data (like UTF-8 for text or Base64 for binary).

type user struct {
	Name     string `json:"name" xml:"name"`
	Age      int    `json:"age" xml:"age"`
	Phone    string `json:"phone" xml:"phone_number"`
	Password string `json:"-" xml:"-"`
	IsActive bool   `json:"active" xml:"active"`
}

func main() {
	fmt.Println("Encoding...")
	u := user{
		Name:  "John Smith",
		Age:   45,
		Phone: "13812345678",
	}

	enc := json.NewEncoder(os.Stdout)
	if err := enc.Encode(&u); err != nil {
		log.Fatal(err)
	}
}
