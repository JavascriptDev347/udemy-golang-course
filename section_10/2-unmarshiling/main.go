package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type user struct {
	Name     string  `json:"name" xml:"name"`
	Email    string  `json:"email" xml:"email"`
	Password string  `json:"-" xml:"-"`
	Age      int     `json:"age" xml:"age"`
	IsActive bool    `json:"isActive" xml:"isActive"`
	Role     string  `json:"-" xml:"role"`
	Profile  profile `json:"profile" xml:"profile"`
}

var payload = `{
 "name": "Jane",
 "age": 20,
 "phone": "123-456-789",
 "active": true,
 "profile": {
  "url": "https://www.jane.co.id"
 }
}
`

type profile struct {
	URL string `json:"url"`
}

func main() {
	fmt.Println("Unmarshiling...")

	var u user
	err := json.Unmarshal([]byte(payload), &u)
	if err != nil {
		fmt.Println(err)
		log.Fatal("Error unmarshaling JSON:", err)
	}

	fmt.Println(u)

	bs, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(bs))
}
