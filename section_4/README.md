# Section 4: Functions, Errors, Defer, and Panic Handling

[Back to main README](../README.md)

This section builds on Go basics and focuses on **functions**, **recursion**, **closures**, **variadic parameters**, **multiple return values**, **custom errors**, **defer**, and **panic/recover**.

The lessons are arranged from simple function usage to more advanced control flow and error handling, so you can learn them in the same order they appear in the folder.

## Table of Contents

1. [1-functions](#1-functions)
2. [2-functions-2](#2-functions-2)
3. [3-variadic-func](#3-variadic-func)
4. [4-multiple-value-func](#4-multiple-value-func)
5. [5-errors](#5-errors)
6. [6-defer](#6-defer)
7. [7-panic-recovery](#7-panic-recovery)
8. [project](#project)

---

## 1-functions

**File:** `1-functions/main.go`

### What this lesson shows
This file introduces the most basic kind of Go functions:
- functions with no return value
- functions with one return value
- functions with multiple parameters
- simple validation before returning a result

### Functions in the file

#### `greet(name string)`
Prints a friendly greeting using `fmt.Printf`.

```go
func greet(name string) {
	fmt.Printf("Hello, %s!\n", name)
}
```

**Lesson:**
- function parameters have explicit types
- you can use formatted output to build messages

#### `add(x int, y int)`
Prints the sum of two integers.

```go
func add(x int, y int) {
	fmt.Printf("%d + %d = %d\n", x, y, x+y)
}
```

**Lesson:**
- functions can do work without returning a value
- parameter names and types make the intention clear

#### `calculateArea(w, h float64) float64`
Returns the area of a rectangle.

```go
func calculateArea(w, h float64) float64 {
	if w < 0 || h < 0 {
		fmt.Println("Width and height must be non-negative.")
		return 0
	}
	return w * h
}
```

**Lesson:**
- functions can validate input before continuing
- early return is a common Go pattern
- the return type appears after the parameter list

### Main takeaways
- Go functions are strongly typed.
- A function may print, return a value, or both.
- Validation before computing helps prevent invalid results.

---

## 2-functions-2

**File:** `2-functions-2/main.go`

### What this lesson shows
This file introduces two important ideas:
- **recursion**
- **closures**

### Functions in the file

#### `factorial(n int) int`
Computes a factorial recursively.

```go
func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}
```

**How it works:**
- when `n <= 1`, the function stops recursion and returns `1`
- otherwise it multiplies `n` by the factorial of `n-1`

**Lesson:**
- recursion means a function calls itself
- every recursive function needs a stopping condition

#### `intSeq() func() int`
Returns a function that remembers state.

```go
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
```

**Lesson:**
- this is a closure
- the inner function captures `i` from the surrounding scope
- each call to the returned function updates the same `i`

### Main takeaways
- recursion is useful for problems that naturally repeat smaller subproblems
- closures let you keep state without using global variables

---

## 3-variadic-func

**File:** `3-variadic-func/main.go`

### What this lesson shows
This file explains **variadic functions** — functions that can accept any number of arguments.

### Functions in the file

#### `sum(numbers ...int) int`
Adds all provided integers.

```go
func sum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}
```

**Lesson:**
- `...int` means “zero or more integers”
- inside the function, `numbers` behaves like a slice
- range loops are a natural fit for variadic input

#### `config(numbers ...int)`
Demonstrates optional input.

```go
func config(numbers ...int) {
	if len(numbers) > 0 {
		first := numbers[0]
		fmt.Println("First number:", first)
	} else {
		fmt.Println("Default value")
	}
}
```

**Lesson:**
- a variadic function can be called with no arguments
- always check length before accessing `numbers[0]`

### Main takeaways
- variadic functions make APIs flexible
- use `len()` to avoid index errors
- a variadic parameter is still type-safe

---

## 4-multiple-value-func

**File:** `4-multiple-value-func/main.go`

### What this lesson shows
This file focuses on **multiple return values**, which are very common in Go.

### Functions in the file

#### `divide(a, b float64) (float64, error)`
Returns the result of division and an error if division is invalid.

```go
func divide(a, b float64) (float64, error) {
	if b == 0 {
		fmt.Println("divide by zero")
		return 0, errors.New("divide by zero")
	}
	return a / b, nil
}
```

**Lesson:**
- Go often returns a normal result and an `error`
- `nil` means no error occurred
- callers should check `err` before trusting the result

#### `splitName(fullName string) (firstName, lastName string)`
Splits a full name into two named return values.

```go
func splitName(fullName string) (firstName, lastName string) {
	parts := strings.Split(fullName, " ")
	firstName = parts[0]
	lastName = parts[1]

	return
}
```

**Lesson:**
- Go supports **named return values**
- the bare `return` works because the named values are already set
- this example assumes the input contains at least two parts

### Important caveat
`splitName` will panic if the name does not contain a space, because it directly reads `parts[1]`.
That makes it a good teaching example, but in real code you should validate the length first.

### Main takeaways
- multiple return values are a core Go style
- `error` is usually the second return value
- named returns can make functions shorter, but they should still be used carefully

---

## 5-errors

**File:** `5-errors/main.go`

### What this lesson shows
This file demonstrates **sentinel errors**, **custom error types**, and **error comparison with `errors.Is`**.

### Variables and types in the file

#### Sentinel errors
```go
var ErrDivisionByZero = errors.New("divide by zero")
var ErrNumTooLarge = errors.New("number too large")
```

**Lesson:**
- sentinel errors are reusable named error values
- they let you check for specific failure cases later

#### `OpError` custom error type
```go
type OpError struct {
	Op      string
	Code    int
	Message string
	Time    time.Time
}
```

This type represents a richer error with metadata.

#### `Error() string`
```go
func (e OpError) Error() string {
	return e.Message
}
```

**Lesson:**
- any type that implements `Error() string` satisfies the `error` interface
- this allows custom error structures

#### `NewOpError(...)`
Creates a pointer to an `OpError`.

#### `DoSomething() error`
Returns a custom error created by `NewOpError`.

### Functions in the file

#### `divide(a, b int) (int, error)`
Checks for invalid input and returns a specific error.

```go
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, ErrDivisionByZero
	}

	if a > 1000 {
		return 0, ErrNumTooLarge
	}

	return a / b, nil
}
```

**Lesson:**
- custom conditions can return specific errors
- returning a named error makes later checking easier

### Main function behavior
```go
value, err := divide(1001, 1)
```
This triggers the `ErrNumTooLarge` branch.

Then:
```go
if errors.Is(err, ErrDivisionByZero) { ... }
else if errors.Is(err, ErrNumTooLarge) { ... }
```
checks which error happened.

### Main takeaways
- use `error` values for expected problems
- use custom error types when you need more details
- use `errors.Is` to compare errors safely

---

## 6-defer

**File:** `6-defer/main.go`

### What this lesson shows
This file explains **defer**, which schedules a function to run when the surrounding function exits.

### Comments in the file
The file includes helpful comments in Uzbek explaining that:
- `defer` runs before the function fully returns
- it is the last action to happen in that function
- it is often used for cleanup

### Function in the file

#### `simpleDefer()`
```go
func simpleDefer() {
	fmt.Println("Simple defer function started")
	defer fmt.Println("This will be printed at the end of the function")
	fmt.Println("Simple defer function ended")
}
```

**Lesson:**
- the `defer` line does not execute immediately
- it runs after the rest of `simpleDefer()` finishes
- deferred calls follow **LIFO** order: last deferred, first executed

### Main function behavior
```go
func main() {
	defer fmt.Println("Main func with defer")
	fmt.Println("Defer")
	simpleDefer()
}
```

**Lesson:**
- `main` also has a deferred print
- the deferred line in `main` runs after `main` is finishing
- this helps demonstrate execution order across nested calls

### Main takeaways
- use defer for cleanup and final actions
- deferred calls are executed even when a function is leaving normally
- `defer` is also triggered during panic unwinding

---

## 7-panic-recovery

**File:** `7-panic-recovery/main.go`

### What this lesson shows
This file demonstrates the relationship between **panic**, **defer**, and **recover**.

### Functions in the file

#### `mightPanic(shouldPanic bool)`
```go
func mightPanic(shouldPanic bool) {
	if shouldPanic {
		panic("something went wrong in mightPanic")
	}

	fmt.Println("This function executed without a panic")
}
```

**Lesson:**
- if `shouldPanic` is true, the function stops immediately
- `panic()` aborts normal flow and starts stack unwinding

#### `recoverable()`
```go
func recoverable() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	mightPanic(true)
}
```

**Lesson:**
- `recover()` only works inside a deferred function
- it catches the panic value
- it prevents the whole program from crashing

#### `main()`
Calls `recoverable()` to demonstrate safe recovery.

### What happens at runtime
1. `main()` calls `recoverable()`
2. `recoverable()` registers a deferred recovery function
3. `mightPanic(true)` triggers a panic
4. Go begins unwinding the stack
5. the deferred function runs
6. `recover()` catches the panic and prints the message
7. the program exits normally

### Main takeaways
- `panic` is for abnormal situations
- `recover` is for handling panics inside deferred functions
- `recover` only works in the same goroutine

---

## project

**File:** `project/main.go`

### What this project shows
This project combines several function and error-handling concepts into one practical example.

### Types and constants

#### `MathError`
```go
type MathError struct {
	Operation string
	InputA    int
	InputB    int
	Message   string
}
```

This is a custom error type that stores:
- the operation name
- the input values
- a readable message

#### Constants
```go
const (
	division       = "Division"
	divisionErrMsg = "division by zero is not allowed"
)
```

### Methods and functions

#### `Error() string`
```go
func (e MathError) Error() string {
	var inputs []string
	if e.Operation == division {
		inputs = append(inputs, fmt.Sprintf("a=%d", e.InputA))
		inputs = append(inputs, fmt.Sprintf("b=%d", e.InputB))
	}

	return fmt.Sprintf("Math error in %s (%s): %s",
		e.Operation,
		strings.Join(inputs, ","),
		e.Message)
}
```

**Lesson:**
- custom error types can format helpful debugging information
- string joining is used to build a readable message

#### `Sum(numbers ...int) int`
```go
func Sum(numbers ...int) int {
	defer fmt.Println("Sum finished")

	total := 0
	for _, n := range numbers {
		total += n
	}
	return total
}
```

**Lesson:**
- variadic parameters are useful in real helper functions
- `defer` can be used for final logging or cleanup

#### `SafeDivision(a, b int) (int, error)`
```go
func SafeDivision(a, b int) (int, error) {
	if b == 0 {
		return 0, &MathError{...}
	}

	return a / b, nil
}
```

**Lesson:**
- this is the safer version of division
- errors are returned instead of causing a panic
- callers can decide how to handle the problem

### Main function behavior
```go
fmt.Println(Sum(1, 2, 3))
value, err := SafeDivision(10, 0)
```

**Lesson:**
- the project demonstrates both success flow and error flow
- it combines variadic functions, custom errors, and `defer`

### Main takeaways
- this is closer to real-world Go code than the earlier lessons
- good function design makes code easier to reuse and test
- custom errors improve debugging and user feedback

---

## How to run the lessons

```bash
cd section_4/1-functions && go run main.go
cd section_4/2-functions-2 && go run main.go
cd section_4/3-variadic-func && go run main.go
cd section_4/4-multiple-value-func && go run main.go
cd section_4/5-errors && go run main.go
cd section_4/6-defer && go run main.go
cd section_4/7-panic-recovery && go run main.go
cd section_4/project && go run main.go
```

---

## Learning path summary

- **Start with functions**: `1-functions`
- **Then learn recursion and closures**: `2-functions-2`
- **Move to variadic input**: `3-variadic-func`
- **Practice returning values and errors**: `4-multiple-value-func`
- **Understand custom error design**: `5-errors`
- **Learn cleanup with defer**: `6-defer`
- **Finish with panic and recovery**: `7-panic-recovery`
- **Apply everything in the project**: `project`

---

## Back to the course home

[Open the main repository README](../README.md)

