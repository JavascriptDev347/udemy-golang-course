package main

import (
	"fmt"
)

func main() {
	fmt.Println("=== RANGE LOOPS IN GO ===\n")

	// ============================================
	// 1. RANGE OVER SLICES (Arrays and Lists)
	// ============================================
	fmt.Println("1. RANGE OVER SLICES")
	fmt.Println("-" + "---------------------------------")

	fruits := []string{"Apple", "Banana", "Cherry", "Date"}

	// Range with index and value
	fmt.Println("\nWith Index and Value:")
	for index, fruit := range fruits {
		fmt.Printf("Index: %d, Name: %s\n", index, fruit)
	}

	// Range with only index (value ignored with _)
	fmt.Println("\nWith Index Only:")
	for index := range fruits {
		fmt.Printf("Index: %d\n", index)
	}

	// Range with only value (use blank identifier to skip index)
	fmt.Println("\nWith Value Only (ignoring index):")
	for _, fruit := range fruits {
		fmt.Printf("Fruit: %s\n", fruit)
	}

	// ============================================
	// 2. RANGE OVER MAPS (Key-Value Pairs)
	// ============================================
	fmt.Println("\n2. RANGE OVER MAPS")
	fmt.Println("-" + "---------------------------------")

	studentGrades := map[string]int{
		"Alice": 95,
		"Bob":   87,
		"Carol": 92,
		"David": 88,
	}

	// Range over map returns key and value
	fmt.Println("\nIterating over Map:")
	for name, grade := range studentGrades {
		fmt.Printf("Student: %s, Grade: %d\n", name, grade)
	}

	// Range with only key
	fmt.Println("\nWith Key Only:")
	for name := range studentGrades {
		fmt.Printf("Student Name: %s\n", name)
	}

	// Range with only value
	fmt.Println("\nWith Value Only:")
	for _, grade := range studentGrades {
		fmt.Printf("Grade: %d\n", grade)
	}

	// ============================================
	// 3. RANGE OVER STRINGS (Characters)
	// ============================================
	fmt.Println("\n3. RANGE OVER STRINGS")
	fmt.Println("-" + "---------------------------------")

	message := "Hello, Go!"

	// Range over string returns rune position and character
	fmt.Println("\nIterating over String (Runes):")
	for index, char := range message {
		fmt.Printf("Position: %d, Character: %c (Unicode: %d)\n", index, char, char)
	}

	// Range with only value
	fmt.Println("\nString Characters Only:")
	for _, char := range message {
		fmt.Printf("Character: %c, ", char)
	}
	fmt.Println()

	// ============================================
	// 4. RANGE OVER NUMBERS (Numeric Range)
	// ============================================
	fmt.Println("\n4. NUMERIC RANGE (Using Simple For Loop)")
	fmt.Println("-" + "---------------------------------")

	// Note: Go doesn't have range(n) like Python
	// We use traditional for loop or libraries
	fmt.Println("\nTraditional Loop (0 to 4):")
	for i := 0; i < 5; i++ {
		fmt.Printf("Number: %d\n", i)
	}

	// ============================================
	// 5. PRACTICAL EXAMPLES
	// ============================================
	fmt.Println("\n5. PRACTICAL EXAMPLES")
	fmt.Println("-" + "---------------------------------")

	// Example 1: Calculate sum of numbers
	fmt.Println("\nSum of Numbers in Slice:")
	numbers := []int{10, 20, 30, 40, 50}
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	fmt.Printf("Sum: %d\n", sum)

	// Example 2: Find specific item in slice
	fmt.Println("\nSearch for Item in Slice:")
	searchFor := "Cherry"
	found := false
	for _, fruit := range fruits {
		if fruit == searchFor {
			found = true
			break // Exit loop when found
		}
	}
	if found {
		fmt.Printf("%s found in fruits list!\n", searchFor)
	} else {
		fmt.Printf("%s NOT found in fruits list!\n", searchFor)
	}

	// Example 3: Create a new slice from existing one
	fmt.Println("\nDouble all Numbers:")
	doubledNumbers := make([]int, len(numbers))
	for index, num := range numbers {
		doubledNumbers[index] = num * 2
	}
	fmt.Printf("Original: %v\n", numbers)
	fmt.Printf("Doubled: %v\n", doubledNumbers)

	// Example 4: Filter items from map
	fmt.Println("\nStudents with Grade >= 90:")
	for name, grade := range studentGrades {
		if grade >= 90 {
			fmt.Printf("%s: %d\n", name, grade)
		}
	}

	// ============================================
	// 6. NESTED RANGE LOOPS
	// ============================================
	fmt.Println("\n6. NESTED RANGE LOOPS")
	fmt.Println("-" + "---------------------------------")

	// Matrix (2D slice)
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println("\nIterating 2D Matrix:")
	for row, rowData := range matrix {
		for col, value := range rowData {
			fmt.Printf("matrix[%d][%d] = %d\n", row, col, value)
		}
	}

	// ============================================
	// 7. BLANK IDENTIFIER USAGE
	// ============================================
	fmt.Println("\n7. BLANK IDENTIFIER (_) FOR UNUSED VALUES")
	fmt.Println("-" + "---------------------------------")

	fmt.Println("\nIgnoring Index (use _):")
	for _, fruit := range fruits {
		fmt.Printf("- %s\n", fruit)
	}

	fmt.Println("\nIgnoring Value (use _):")
	for i := range fruits {
		fmt.Printf("Index: %d\n", i)
	}

	// ============================================
	// 8. COMMON PATTERNS & BEST PRACTICES
	// ============================================
	fmt.Println("\n8. COMMON PATTERNS")
	fmt.Println("-" + "---------------------------------")

	// Pattern 1: Counting occurrences
	fmt.Println("\nCount Specific Grade:")
	gradeCount := 0
	targetGrade := 90
	for _, grade := range studentGrades {
		if grade >= targetGrade {
			gradeCount++
		}
	}
	fmt.Printf("Students with grade >= %d: %d\n", targetGrade, gradeCount)

	// Pattern 2: Building a new slice
	fmt.Println("\nFilter Numbers (only > 25):")
	largNumbers := []int{}
	for _, num := range numbers {
		if num > 25 {
			largNumbers = append(largNumbers, num)
		}
	}
	fmt.Printf("Original: %v\n", numbers)
	fmt.Printf("Filtered: %v\n", largNumbers)

	// Pattern 3: Index-based access
	fmt.Println("\nAccessing by Index within Range:")
	for idx := range fruits {
		if idx%2 == 0 {
			fmt.Printf("Even index %d: %s\n", idx, fruits[idx])
		}
	}
}

