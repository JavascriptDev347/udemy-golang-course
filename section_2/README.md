# Section 2: Control Flow & Loops in Go

This section covers essential control flow structures in Go programming language, including loops, conditional statements, and switch cases.

## Table of Contents
1. [For Loops](#for-loops)
2. [If-Else Statements](#if-else-statements)
3. [Switch Statements](#switch-statements)
4. [Range Loops](#range-loops) ⭐ NEW
5. [Break & Continue](#break--continue-statements) ⭐ NEW
6. [Logical Operators](#logical-operators) ⭐ NEW
7. [Practical Project](#practical-project)

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

## Range Loops

**File:** `4-range-loops/main.go`

### Overview
Range loops are the most common and idiomatic way to iterate in Go. The `range` keyword works with slices, arrays, maps, strings, and channels, returning the index/key and value for each iteration.

### Key Concepts:
- **Range over slices**: Iterate through array elements with index and value
- **Range over maps**: Iterate through key-value pairs
- **Range over strings**: Iterate through characters (runes)
- **Blank identifier (_)**: Skip unwanted values (index or value)
- **Index only**: Use range to get just the position
- **Value only**: Ignore index, process only the value

### Example Code:
```go
// Range over slice
fruits := []string{"Apple", "Banana", "Cherry"}
for index, fruit := range fruits {
    fmt.Printf("Index: %d, Name: %s\n", index, fruit)
}

// Range over map
students := map[string]int{
    "Alice": 95,
    "Bob":   87,
}
for name, grade := range students {
    fmt.Printf("Student: %s, Grade: %d\n", name, grade)
}

// Range over string
message := "Hello"
for index, char := range message {
    fmt.Printf("Position: %d, Character: %c\n", index, char)
}

// Using blank identifier to skip index
for _, fruit := range fruits {
    fmt.Printf("Fruit: %s\n", fruit)
}
```

### Learning Points:
- Range is the idiom ic way to iterate in Go
- Always returns two values: index/key and element/value
- Use underscore `_` to ignore unwanted values
- Range over strings gives you rune positions and characters
- Range over maps returns keys and values in random order
- Most performant and readable looping method

---

## Break & Continue Statements

**File:** `5-break-continue/main.go`

### Overview
Break and continue statements give you fine-grained control over loop execution. They allow you to exit loops early or skip iterations based on conditions.

### Key Concepts:
- **Break**: Immediately exit/terminate the current loop
- **Continue**: Skip the current iteration and go to the next one
- **Labeled break**: Exit from outer loops (advanced)
- **Early termination**: Stop processing when condition is met
- **Conditional skipping**: Skip iterations that don't meet criteria

### Example Code:
```go
// Break - exit loop when condition is met
for i := 1; i <= 10; i++ {
    if i == 5 {
        fmt.Println("Breaking at 5")
        break // Exit loop immediately
    }
    fmt.Printf("%d ", i)
}

// Continue - skip to next iteration
for i := 1; i <= 10; i++ {
    if i%2 == 0 {
        continue // Skip even numbers
    }
    fmt.Printf("%d ", i) // Only prints odd numbers
}

// Labeled break (advanced)
OuterLoop:
for i := 1; i <= 3; i++ {
    for j := 1; j <= 3; j++ {
        if i == 2 && j == 2 {
            break OuterLoop // Breaks out of outer loop
        }
    }
}
```

### Learning Points:
- Break exits only the innermost loop
- Continue only affects the current loop iteration
- Labeled break can exit outer loops (use sparingly)
- Combine with conditions for powerful flow control
- Essential for early termination of searches or validations
- Helps write efficient and readable code

---

## Logical Operators

**File:** `6-logical-operators/main.go`

### Overview
Logical operators combine multiple boolean conditions, allowing you to create complex decision-making logic. Go supports AND, OR, and NOT operators with important optimization features.

### Key Concepts:
- **AND (&&)**: Both conditions must be TRUE
- **OR (||)**: At least one condition must be TRUE
- **NOT (!)**: Negates/inverts a boolean value
- **Short-circuit evaluation**: Optimizer skips unnecessary evaluations
- **Operator precedence**: ! > && > ||
- **De Morgan's Laws**: Mathematical principles for logical equivalence

### Example Code:
```go
// AND operator - both conditions must be true
age := 25
income := 50000
if age >= 21 && income >= 30000 {
    fmt.Println("Eligible for loan")
}

// OR operator - at least one condition must be true
day := "Saturday"
if day == "Saturday" || day == "Sunday" {
    fmt.Println("It's weekend!")
}

// NOT operator - negates the condition
isRaining := true
if !isRaining {
    fmt.Println("It's not raining, go outside!")
}

// Combining operators (use parentheses for clarity)
score := 85
attendance := 95
if (score >= 80 && attendance >= 90) || hasExcuse {
    fmt.Println("Eligible for exception")
}
```

### Learning Points:
- AND (&&) has higher precedence than OR (||) - can cause unexpected results without parentheses
- Use parentheses to make complex conditions clear and maintainable
- Short-circuit evaluation improves performance (AND stops if first is false, OR stops if first is true)
- De Morgan's Laws can simplify complex boolean logic
- Avoid overly complex nested conditions - break them into multiple if statements instead
- Use meaningful variable names to make conditions self-documenting

---

## Practical Project

**File:** `project/main.go`

### Overview
This real-world project demonstrates how to combine all the control flow concepts learned in this section into a practical application.

### What It Does:
- Manages product inventory with prices and sales
- Calculates order totals
- Applies discounts to sale items
- Validates product existence
- Uses maps, loops, conditionals, and functions together

### Key Techniques Used:
- Maps for data storage
- Range loops for iteration
- Conditional logic with if-else
- String manipulation
- Functions with error handling (ok idiom)
- Real-world calculations

### Learning Points:
- Combining multiple concepts in a single application
- Practical error handling patterns
- Working with Go's standard library (strings package)
- Realistic business logic implementation
- Clean code organization with helper functions

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

cd section_2/4-range-loops
go run main.go

cd section_2/5-break-continue
go run main.go

cd section_2/6-logical-operators
go run main.go

cd section_2/project
go run main.go
```

---

## Summary

This comprehensive section covers all fundamental control flow structures in Go:

**Basic Loops:**
- **For Loops** - C-style, while-style, and infinite loops
- **Range Loops** - Idiomatic iteration over collections ⭐ NEW

**Decision Making:**
- **If-Else** - Basic and chained conditionals
- **Switch** - Multi-way branching
- **Logical Operators** - Complex conditions with AND, OR, NOT ⭐ NEW

**Flow Control:**
- **Break & Continue** - Loop control and optimization ⭐ NEW

**Real-World Application:**
- **Practical Project** - Combining all concepts

Mastering these constructs is essential for writing effective and idiomatic Go programs!

---

## Progression Chart

```
Beginner Level:
├── 1-for-loop        (C-style & while-style loops)
├── 2-if-else         (Basic conditionals)
└── 3-switch          (Multi-way branching)

Intermediate Level:
├── 4-range-loops     (Idiomatic iteration) ⭐ NEW
├── 5-break-continue  (Flow control) ⭐ NEW
├── 6-logical-operators (Complex conditions) ⭐ NEW
└── project           (Practical application)
```

## Related Resources
- [Go For Loop Documentation](https://golang.org/doc/effective_go#for)
- [Go If Statements Documentation](https://golang.org/doc/effective_go#if)
- [Go Switch Statements Documentation](https://golang.org/doc/effective_go#switch)
- [Go By Example - Range](https://gobyexample.com/range)
- [Effective Go](https://golang.org/doc/effective_go)

## Practice Tips

1. **Run the code**: Don't just read - run each example and modify it
2. **Combine concepts**: Try using range loops with conditionals
3. **Avoid nested complexity**: Break complex logic into variables
4. **Use break/continue wisely**: Don't overuse labeled breaks
5. **Clear conditions**: Always use parentheses in complex logical expressions

---

**Next Steps:** After mastering control flow, explore functions, error handling, and data structures in the next sections!

