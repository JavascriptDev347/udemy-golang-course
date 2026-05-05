package main

import (
	"fmt"
	"strings"
)

func main() {
	// Strings are immutable sequences of bytes that represent text.
	//They are a fundamental data type in Go and are used to store and manipulate textual data.

	// clone func from strings package
	s1 := "Hello russel"
	s2 := strings.Clone(s1)
	fmt.Println(s2 == s1)

	// compare func from strings package
	c1 := "abs"
	c2 := "bdc"
	fmt.Println(strings.Compare(c1, c2)) // -1 because c1 < c2

	// builder func from strings package
	b := strings.Builder{}
	b.Write([]byte("Hello world"))
	fmt.Println(b.String())

	// lower upper trim funcs from strings package
	fmt.Println(strings.ToLower(s1))
	fmt.Println(strings.ToUpper(s2))
	fmt.Println(strings.TrimSpace("   da  "))

	fmt.Println(strings.HasSuffix("test@gmail.com", "gmail.com"))
	fmt.Println(strings.HasPrefix("test@gmail.com", "test"))
	fmt.Println(strings.Replace("test@gmail.com", "test", "john", 1))

	parts := strings.Split("test@gmail.com", "@")
	username, domain := parts[0], parts[1]
	fmt.Println(username, domain)

	parts = strings.Fields("jane example.com")
	username, domain = parts[0], parts[1]
	fmt.Println(username, domain)

	fmt.Println(strings.Join(parts, ","))

}
