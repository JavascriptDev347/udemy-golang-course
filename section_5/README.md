# Section 5: Structs, Methods, Interfaces & Generics

This section teaches how to model data with structs, attach behavior with methods, define and use interfaces, and write generic functions in Go.

Contents (lesson files)

- `1-struct/main.go` — Basic structs and constructors
- `2-methods/main.go` — Methods, pointer vs value receivers, mutability
- `3-interfaces/main.go` — Interfaces and polymorphism
- `4-generics/main.go` — Simple generics example (type parameters)

---

## 1-struct

**File:** `1-struct/main.go`

What it demonstrates
- Declaring a `struct` type: `Employee` with fields like `ID`, `FirstName`, `LastName`, `Position`, `Salary`, `IsActive`, `JoinedAt`.
- Creating struct values using composite literals and a constructor-like function `NewEmployee`.
- Printing struct values with `fmt.Printf("%+v\n", value)` to show field names and values.

Key concepts
- Structs are Go's way to group related fields into one composite value.
- You can create helper functions (constructors) to initialize structs consistently.
- Struct fields are exported (capitalized) or unexported (lowercase) which matters across packages.

Run:

```bash
cd section_5/1-struct && go run main.go
```

---

## 2-methods

**File:** `2-methods/main.go`

What it demonstrates
- Defining methods on types using receiver syntax: `func (e *Employee) FullName() string` and `func (e *Employee) EditActive(active bool) bool`.
- Pointer receiver (`*Employee`) allows the method to modify the receiver's fields.
- Nil-check inside method (`if e == nil`) to make methods safer when called on a nil pointer.
- Appending struct values to slices and how pointer vs value semantics affect modifications.

Key concepts
- Value receiver (e Employee) receives a copy of the value — changes inside the method do not affect the caller's copy.
- Pointer receiver (e *Employee) receives an address — methods can modify the original value.
- Use pointer receivers when the method needs to modify state or when copying the value is expensive.

Run:

```bash
cd section_5/2-methods && go run main.go
```

Notes and recommendations
- In this file `FullName` uses a pointer receiver and checks for `nil` to avoid panics.
- `EditActive` demonstrates mutating a field through a pointer receiver; that change is visible to callers.

---

## 3-interfaces

**File:** `3-interfaces/main.go`

What it demonstrates
- Defining an interface `Person` with methods `GetName() string` and `GetID() int`.
- Two concrete types (`Employee` and `BusinessEmployee`) that implement `GetName()` (both) and are expected to implement `GetID()`.
- A `displayName(p Person)` function that accepts any value satisfying the `Person` interface — demonstrating polymorphism.

Important note about this lesson
- The `GetID()` methods for both `Employee` and `BusinessEmployee` are currently left as TODOs and call `panic("implement me")`.
  - To make the example runnable, implement these methods to return the `ID` field, e.g.:

```go
func (e Employee) GetID() int { return e.ID }
func (b BusinessEmployee) GetID() int { return b.ID }
```

Key concepts
- A type implements an interface implicitly by providing the required methods — no `implements` keyword.
- Interfaces let code operate on abstract behavior rather than concrete types.
- Small, focused interfaces (single-method) are recommended in idiomatic Go.

Run (after implementing `GetID`):

```bash
cd section_5/3-interfaces && go run main.go
```

---

## 4-generics

**File:** `4-generics/main.go`

What it demonstrates
- A generic `Sum` function using type parameters:

```go
func Sum[T int | float32 | float64](numbers ...T) T
```

- The function accepts a variadic list of a constrained set of numeric types and returns their sum of the same type.
- Prints the deduced type at runtime using `fmt.Printf("%T\n", value)`.

Key concepts
- Generics (type parameters) enable writing reusable code that works with multiple types while preserving type safety.
- The example uses a union type constraint (`int | float32 | float64`) which restricts allowed argument types.

Run:

```bash
cd section_5/4-generics && go run main.go
```

---

## Overall takeaways & suggested improvements

- Structs and methods are the foundation of Go's compositional design; prefer small types with clear responsibilities.
- Decide between pointer and value receivers based on whether the method mutates the receiver or if the value is expensive to copy.
- Implement the missing `GetID()` methods in `3-interfaces` to make that example complete and runnable.
- The generics example is simple and effective; consider adding more examples (e.g., a `MapKeys` generic helper, or generic stack/queue) to show additional patterns.

## How to run all section examples

```bash
cd section_5/1-struct && go run main.go
cd ../2-methods && go run main.go
cd ../3-interfaces && go run main.go   # implement GetID first
cd ../4-generics && go run main.go
```

## Further reading

- Effective Go: https://golang.org/doc/effective_go
- Go blog on generics (proposal and design): https://go.dev/doc/go1.18
- Go by Example (structs, interfaces): https://gobyexample.com/

---

If you want, I can also implement the missing `GetID()` methods in `section_5/3-interfaces/main.go` and run the examples to verify everything compiles. 
