# Section 2: Control Flow & Loops in Go

This section covers essential control flow structures in Go programming language, including loops, conditional statements, and switch cases.

## Table of Contents
1. [For Loops](#for-loops)
2. [If-Else Statements](#if-else-statements)
3. [Switch Statements](#switch-statements)

---

## For Loops

**File:** `1-foor-loop/main.go`

### Overview
Go supports `for` loops which are the primary looping construct. Unlike other languages, Go doesn't have `while` or `do-while` loops - the `for` statement covers all looping needs.

### Key Concepts:
- **C-style for loop**: `for i := 1; i < 10; i++`
- **While-style loop**: `for condition { ... }`
- **Range loops**: Used for iterating over collections
- **Infinite loops**: `for { ... }`

### Example Code:
```go
// C-style for loop with initialization, condition, and increment
a := 10
for i := 1; i < 10; i++ {
    a *= i
    println(a)
}

// While-style loop (condition without initialization and increment)
k := 3
for k > 0 {
    println(k)
    k--
}
```

### Learning Points:
- Use `for` keyword with three optional parts: `init; condition; post`
- Loop counter auto-increments with `i++`
- Short variable declaration with `:=` inside loop scope
- Break and continue statements can be used to control loop flow

---

## If-Else Statements

**File:** `2-if-else/main.go`

### Overview
Conditional statements allow the program to make decisions based on conditions. Go provides `if`, `else if`, and `else` keywords for branching logic.

### Key Concepts:
- **Simple if statement**: Check a single condition
- **if-else statement**: Execute different code for true/false conditions
- **if-else if-else chain**: Check multiple conditions sequentially
- **Assignments in conditions**: You can initialize variables directly in the if statement
- **Map access with ok idiom**: Check if a key exists while retrieving its value

### Example Code:
```go
// Simple if-else
tmp := 25
if tmp > 30 {
    fmt.Println("HOT")
} else {
    fmt.Println("NORM")
}

// If-else if-else chain
score := 87
if score >= 90 {
    fmt.Println("A")
} else if score >= 80 {
    fmt.Println("B")
} else if score >= 70 {
    fmt.Println("C")
} else {
    fmt.Println("Failed")
}

// Map access with ok idiom
userAccess := map[string]bool{
    "Russel": true,
    "James":  true,
    "Goose":  false,
}

if hasAccess, ok := userAccess["Russel"]; ok {
    if hasAccess {
        fmt.Println(hasAccess, ok)
    }
}
```

### Learning Points:
- Use curly braces `{}` to define code blocks (they are required)
- Semicolon is NOT needed at the end of lines in Go
- The "ok" idiom is Go's way to check if a key exists in a map
- Multiple conditions can be chained together
- Nested if statements are allowed

---

## Switch Statements

**File:** `3-switch/main.go`

### Overview
Switch statements are used when you have multiple possible values for a single expression. They are cleaner and more readable than long if-else chains. Go's switch has some unique features like fallthrough behavior that's different from other languages.

### Key Concepts:
- **Value switch**: Compare a value against multiple cases
- **Multiple cases**: Use comma to combine multiple values in one case
- **Expression switch**: Evaluate expressions in each case
- **Type switch**: Check the type of a variable using `v := i.(type)`
- **Default case**: Execute if no cases match
- **No fallthrough by default**: Cases don't automatically fall to the next case

### Example Code:
```go
// Basic switch with string value
today := "Sunday"
switch today {
case "Sunday":
    fmt.Println("Sunday day offf")
case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
    fmt.Println("Working day again :cry")
default:
    fmt.Println("Saturday")
}

// Expression switch (evaluates expressions in cases)
switch hour := time.Now().Hour(); {
case hour < 12:
    fmt.Println("Good morning")
case hour < 17:
    fmt.Println("Good afternoon")
default:
    fmt.Println("Good evening")
}

// Type switch (checks the type of value)
checkType := func(i interface{}) {
    switch v := i.(type) {
    case int:
        fmt.Printf("Twice %v is %v\n", v, v*2)
    case string:
        fmt.Printf("%q is a string\n", v)
    case bool:
        fmt.Printf("Twice %v is a bool\n", v)
    default:
        fmt.Printf("I don't know about type %T!\n", v)
    }
}

checkType(2)
checkType("Hello")
```

### Learning Points:
- Cases are evaluated from top to bottom and stop when a match is found
- Multiple values can be checked in a single case using commas
- Expression switch allows you to put an expression after the switch keyword
- Type switch uses the syntax `v := i.(type)` to extract the type and value
- No fallthrough by default (unlike C or Java) - each case must explicitly use `fallthrough` keyword
- Default case is optional and doesn't have to be the last case

---

## How to Run

To run any of these examples:

```bash
cd section_2/<lesson-folder>
go run main.go
```

For example:
```bash
cd section_2/1-foor-loop
go run main.go
```

---

## Summary

This section introduces the fundamental control flow structures in Go:
- **Loops** help repeat code blocks
- **Conditionals (if-else)** allow decision making
- **Switches** provide elegant multi-way branching

Mastering these constructs is essential for writing effective Go programs!

---

## Related Resources
- [Go For Loop Documentation](https://golang.org/doc/effective_go#for)
- [Go If Statements Documentation](https://golang.org/doc/effective_go#if)
- [Go Switch Statements Documentation](https://golang.org/doc/effective_go#switch)

