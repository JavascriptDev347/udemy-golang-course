# Section 3: Data Structures in Go

This section covers essential data structures in Go programming language, including arrays, slices, maps, and pointers. You'll learn how to work with collections and manage memory efficiently.

## Table of Contents
1. [Arrays](#arrays)
2. [Slices](#slices)
3. [Maps](#maps)
4. [Pointers](#pointers)
5. [Advanced Slices](#advanced-slices)
6. [Practical Project](#practical-project)

---

## Arrays

**File:** `1-array/main.go`

### Overview
Arrays are fixed-size collections of elements of the same type. Once declared, the size of an array cannot be changed. Arrays are useful when you know exactly how many elements you need.

### Key Concepts:
- **Fixed size**: Arrays have a fixed length specified at declaration
- **Initialization**: Arrays can be declared empty or with values
- **Indexing**: Elements accessed by index (0-based)
- **Length**: Use `len()` to get array size
- **Iteration**: Use for loops to traverse elements

### Example Code:
```go
// Declare empty array with 2 int elements (defaults to 0)
var numbers [2]int
fmt.Printf("Array elements :%v\n Array length: %v\n", numbers, len(numbers))

// Declare and initialize array with values
primes := [4]int{2, 3, 5, 7}
fmt.Printf("Array elements :%v\n Array length: %v\n", primes, len(primes))

// Iterate through array
for i := 0; i < len(primes); i++ {
    fmt.Printf("Element at index %v is %v\n", i, primes[i])
}
```

### Learning Points:
- Arrays are value types (copying creates a full copy)
- Size is part of the type (e.g., `[4]int` and `[5]int` are different types)
- Rarely used directly in Go; slices are preferred for collections
- Used for fixed-size buffers in performance-critical code

### When to Use Arrays:
- When size is guaranteed to never change
- For performance-critical fixed-size buffers
- Most Go code prefers slices over arrays

---

## Slices

**File:** `2-slices/main.go`

### Overview
Slices are dynamic arrays backed by an underlying array. They're more flexible than arrays as they can grow and shrink. Slices are the primary collection type in Go.

### Key Concepts:
- **Dynamic size**: Slices can grow and shrink at runtime
- **Slice literals**: `[]string{}` creates an empty slice
- **make() function**: Create slices with specific length and capacity
- **Length vs Capacity**: Length is current size, capacity is max before reallocation
- **Flexible**: Best for most collection scenarios

### Example Code:
```go
// Slice literal with initial values
names := []string{
    "John Smith",
    "Jane Smith",
    "John Smith",
}
fmt.Println(names)

// Create slice with make(type, length, capacity)
items := make([]int, 3, 10)
// Creates slice with length 3 and capacity 10

// Empty slice
sliceWithoutMake := []int{}
fmt.Println(sliceWithoutMake)
fmt.Println(items)
```

### Learning Points:
- Slices are reference types (passing by value shares the underlying array)
- `len()` returns the number of elements in the slice
- `cap()` returns the total capacity available
- Slices are preferred over arrays in Go
- Use `make()` when you know desired length and capacity

### Slice Declaration Methods:
1. **Slice literal**: `[]int{1, 2, 3}`
2. **From array**: `array[start:end]`
3. **Using make()**: `make([]int, length, capacity)`

---

## Maps

**File:** `3-maps/main.go`

### Overview
Maps are unordered collections of key-value pairs. They're similar to dictionaries in Python or objects in JavaScript. Each key is unique and maps to a value.

### Key Concepts:
- **Key-value pairs**: Store data with descriptive keys
- **Unordered**: Keys are not stored in order
- **Unique keys**: Each key can only appear once
- **"ok" idiom**: Check if key exists while retrieving value
- **Delete operation**: Remove key-value pairs from map
- **Make function**: Create empty maps with `make(map[keyType]valueType)`

### Example Code:
```go
// Map with initial values
studentGrades := map[string]int{
    "Jane":   62,
    "John":   86,
    "Alice":  89,
    "Russel": 98,
}
fmt.Println("Original map", studentGrades)

// Update existing key
studentGrades["Alice"] = 90

// Add new key
studentGrades["Bob"] = 75

// Check if key exists (ok idiom)
alice, ok := studentGrades["Alice"]
if ok {
    fmt.Println("Topildi: ", alice)
}

// Delete key
delete(studentGrades, "John")
fmt.Println("After delete:", studentGrades)

// Create empty map with make
configs := make(map[string]string)
```

### Learning Points:
- Maps are reference types
- Use the "ok" idiom to check existence and retrieve value safely
- `delete()` function removes a key-value pair
- Iteration order over maps is random
- Maps cannot be nil if declared without make (but will panic on access)
- Very useful for lookups and caching

### Common Operations:
- **Access**: `map[key]` returns value or zero value
- **Update**: `map[key] = newValue`
- **Delete**: `delete(map, key)`
- **Check**: `value, ok := map[key]`

---

## Pointers

**File:** `4-pointers/main.go`

### Overview
Pointers hold memory addresses of variables. They allow you to pass variables by reference and modify their values in functions. Understanding pointers is crucial for efficient Go programming.

### Key Concepts:
- **Address operator (&)**: Get the memory address of a variable
- **Dereference operator (*)**: Access the value at a pointer address
- **Pass by value vs reference**: Modify function parameters
- **Function parameters**: Use pointers to modify variables outside function scope
- **Nil pointers**: Initialize pointers before use

### Example Code:
```go
// Function that doesn't modify original (pass by value)
func modifyValue(val int) {
    val = val * 10
    fmt.Printf("Modified value: %d\n", val)
}

// Function that modifies original (pass by reference with pointer)
func modifyPointer(val *int) {
    if val == nil {
        fmt.Printf("Value: %d\n", val)
        return
    }
    *val = *val * 10
    fmt.Printf("Modified pointer value: %d\n", *val)
}

func main() {
    // Get address of variable
    age := 10
    agePtr := &age
    fmt.Println("Age address: ", &age)
    fmt.Println("Age address: ", agePtr)

    // Pass by value (doesn't modify original)
    val := 10
    modifyValue(val)
    fmt.Println("Original value: ", val) // Still 10

    // Pass by reference (modifies original)
    val2 := 12
    modifyPointer(&val2)
    fmt.Println("Original value: ", val2) // Now 120
}
```

### Learning Points:
- `&variable` gets the address of a variable
- `*pointerVariable` dereferences to access the value
- Pointers allow functions to modify variables in the caller's scope
- Always check for nil before dereferencing
- Pointers are useful for large structured data (avoid copying)
- Go doesn't have pointer arithmetic like C

### When to Use Pointers:
- Modifying variables in function parameters
- Working with large structs (avoid expensive copying)
- Creating reference relationships between data
- Implementing linked lists, trees, etc.

---

## Advanced Slices

**File:** `5-slices-2/main.go`

### Overview
Advanced slice operations include slicing ranges, understanding capacity, and using the `slices` package for common operations like searching and inserting.

### Key Concepts:
- **Slice ranges**: `slice[start:end]` creates a view of the original
- **Length and capacity**: `len()` and `cap()` give different information
- **Slices package**: Go 1.18+ provides helper functions for slices
- **Contains**: Check if slice contains a value
- **Insert**: Add elements at specific positions
- **Full slice expressions**: Control start, end, and capacity

### Example Code:
```go
original := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
fmt.Printf("Original %v, Len %d, cap %d\n", original, len(original), cap(original))

// Slice from start to index 5
s1 := original[:5]
fmt.Printf("S1 %v, Len %d, cap %d\n", s1, len(s1), cap(s1))

// Slice from index 5 to 7
s2 := original[5:7]
fmt.Printf("S2 %v, Len %d, cap %d\n", s2, len(s2), cap(s2))

// Slice from index 7 to end
s3 := original[7:]
fmt.Printf("S3 %v, Len %d, cap %d\n", s3, len(s3), cap(s3))

// Full slice
s4 := original[:]
fmt.Printf("S4 %v, Len %d, cap %d\n", s4, len(s4), cap(s4))

// Check if contains and insert using slices package
slices.Contains(s4, 8)
slices.Insert(s4, 5, 42)
fmt.Printf("Modified S4 %v, Len %d, cap %d\n", s4, len(s4), cap(s4))
```

### Learning Points:
- Slicing creates a view, not a copy (shares underlying array)
- `len()` is the slice length, `cap()` is the total available capacity
- Modifying a slice element affects the underlying array
- `slices` package (Go 1.18+) provides useful utilities
- Slice ranges: `[start:end]` excludes the end index
- Omitting indices: `[:5]`, `[5:]`, `[:]` are valid

### Slice Operations:
- **Create view**: `slice[start:end]`
- **Get length**: `len(slice)`
- **Get capacity**: `cap(slice)`
- **Check contains**: `slices.Contains(slice, value)`
- **Insert element**: `slices.Insert(slice, index, values...)`

---

## Practical Project

**File:** `project/main.go`

### Overview
A contact management system that demonstrates how to combine structs, slices, and maps in a real-world application. It manages a list of contacts with add, find, and list operations.

### What It Does:
- Defines a `Contact` struct with ID, Name, Email, and Phone
- Maintains a slice of contacts (`contactList`)
- Uses a map for fast name-based lookups (`contactIndexByName`)
- Provides functions to add, find, and list contacts
- Prevents duplicate contacts by name

### Key Techniques Used:
```go
// Struct definition
type Contact struct {
    ID    int
    Name  string
    Email string
    Phone string
}

// Global data structures
var contactList []Contact
var contactIndexByName map[string]int
var nextID = 1

// Initialization function
func init() {
    contactList = make([]Contact, 0)
    contactIndexByName = make(map[string]int)
}

// Add contact with duplicate check
func addContact(name, email, phone string) { ... }

// Find contact by name (return pointer)
func findContact(name string) *Contact { ... }

// List all contacts
func ListContacts() { ... }
```

### Learning Points:
- Structs organize related data
- Slices store ordered collections
- Maps provide fast lookups
- Pointers return references to struct elements
- Combining data structures solves real problems
- `init()` function runs automatically on program start
- The `ok` idiom prevents map errors

### Design Patterns:
- **Unique ID generation**: Using a counter for each new record
- **Duplicate prevention**: Check map before adding
- **Fast lookup**: Map indexes into the slice
- **Data integrity**: Return pointers instead of copies

---

## How to Run

To run any of these examples:

```bash
cd section_3/<lesson-folder>
go run main.go
```

For example:
```bash
cd section_3/1-array
go run main.go

cd section_3/2-slices
go run main.go

cd section_3/3-maps
go run main.go

cd section_3/4-pointers
go run main.go

cd section_3/5-slices-2
go run main.go

cd section_3/project
go run main.go
```

---

## Comparison: Arrays vs Slices

| Feature | Arrays | Slices |
|---------|--------|--------|
| Size | Fixed | Dynamic |
| Declaration | `[size]Type` | `[]Type` |
| Memory | Stack (usually) | Heap |
| Copying | Full copy | Reference copy |
| Preferred | Rarely | Most cases |
| Append | Not possible | Yes, with `append()` |

---

## Comparison: Maps vs Slices

| Operation | Maps | Slices |
|-----------|------|--------|
| Ordered | No | Yes |
| Fast lookup by key | Yes | No |
| Iteration | Unordered | Ordered |
| Index type | Any comparable | Integer |
| Nil handling | Panic if nil | Safe to use |
| Use case | Lookups | Collections |

---

## Summary

This section covers the fundamental data structures in Go:

**Collections:**
- **Arrays** - Fixed-size collections (rarely used directly)
- **Slices** - Dynamic, ordered collections (most common)
- **Maps** - Key-value pairs for fast lookups

**Memory Management:**
- **Pointers** - References to memory locations
- Understanding pass-by-value vs pass-by-reference

**Real-World Application:**
- **Practical Project** - Contact management system combining all concepts

Mastering these data structures is essential for writing effective Go programs!

---

## Progression Path

```
Beginner Level:
├── 1-array          (Fixed-size collections)
└── 2-slices         (Dynamic collections)

Intermediate Level:
├── 3-maps           (Key-value storage)
├── 4-pointers       (Memory references)
├── 5-slices-2       (Advanced slicing)
└── project          (Real-world application)
```

---

## Common Mistakes & Best Practices

### Arrays
- ❌ Trying to change array size (impossible)
- ✅ Use slices for dynamic collections instead

### Slices
- ❌ Not checking length before accessing by index
- ✅ Always validate index ranges or use range loops

### Maps
- ❌ Accessing non-existent key without checking
- ✅ Always use the "ok" idiom: `value, ok := map[key]`

### Pointers
- ❌ Dereferencing nil pointers (causes panic)
- ✅ Always check for nil before dereferencing

### Slicing
- ❌ Forgetting that slices share underlying array
- ✅ Be aware of side effects when modifying slices

---

## Related Resources

- [Go Arrays Documentation](https://golang.org/doc/effective_go#arrays)
- [Go Slices Documentation](https://golang.org/doc/effective_go#slices)
- [Go Maps Documentation](https://golang.org/doc/effective_go#maps)
- [Go Pointers Documentation](https://golang.org/doc/effective_go#pointers)
- [Go Slices Package](https://pkg.go.dev/slices)
- [Effective Go](https://golang.org/doc/effective_go)

---

## Practice Tips

1. **Run the code**: Don't just read - execute each example
2. **Modify examples**: Change values and see what happens
3. **Combine concepts**: Use slices with maps, arrays with pointers
4. **Use range loops**: Preferred over index-based loops
5. **Check for nil**: Always validate pointers before use
6. **Leverage maps**: Use them for O(1) lookups instead of searching slices

---

**Next Steps**: After mastering data structures, explore functions, error handling, and concurrent programming in the next sections!

