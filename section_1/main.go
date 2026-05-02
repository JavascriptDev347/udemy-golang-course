package main

import "fmt"

type LogLevel int

const (
	LevelTrace LogLevel = iota
	LevelDebug
	LevelInfo
	LevelWarn
	LevelError
)

var LevelNames = []string{"TRACE", "DEBUG", "INFO", "WARN", "ERROR"}

func (l LogLevel) String() string {
	if l < LevelTrace || l > LevelError {
		return "Unknown"
	}
	return LevelNames[l]
}

func printLogLevel(level LogLevel) {
	fmt.Printf("%s\t %s\n", level, level.String())
}

func main() {
	printLogLevel(LevelTrace)
	printLogLevel(LevelDebug)
	printLogLevel(LevelInfo)
	printLogLevel(LevelWarn)
	printLogLevel(LevelError)
	printLogLevel(123)
}
