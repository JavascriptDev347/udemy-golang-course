package main

import (
	"errors"
	"fmt"
)

type ConfigItem struct {
	Key   string
	Value interface{}
	IsSet bool
}

/*
%v ->the default formatting
%+v -> the default formatting but include field names
%#v -> Go-syntax representation of the value
%T -> the type of the value
%s -> the string representation of the value (if it implements Stringer interface)
%d -> the decimal representation of an integer
%f (%.2f) -> the decimal representation of a floating-point number
%t -> the boolean representation of a value
*/
func (c ConfigItem) String() string {
	return fmt.Sprintf("Key: %s, Value: %s, IsSet: %t", c.Key, c.Value, c.IsSet)
}

func main() {
	appName := "Smart House"
	version := "1.0.1"
	port := 9090
	isEnabled := true

	status := fmt.Sprintf("Application: %s, Version: %s, Port: %d, Enabled: %t", appName, version, port, isEnabled)
	fmt.Println(status)

	item1 := ConfigItem{Key: "API_URL", Value: "http://localhost", IsSet: true}
	item2 := ConfigItem{Key: "TIMEOUT", Value: 5000, IsSet: false}
	item3 := ConfigItem{Key: "DEBUG_MODE", Value: false, IsSet: true}
	fmt.Printf("Config Item 1 (%%v): %v\n", item1)
	fmt.Printf("Config Item 2 (%%+v): %+v\n", item2)
	fmt.Printf("Config Item 3 (%%#v): %#v\n", item3)

	err := errors.New("test")

	fmt.Errorf("here is the error on port %d: %w", port, err)
}
