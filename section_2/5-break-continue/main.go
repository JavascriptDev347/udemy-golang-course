package main

import (
	"fmt"
)

func main() {
	fmt.Println("=== BREAK & CONTINUE STATEMENTS ===\n")

	// ============================================
	// 1. BREAK STATEMENT
	// ============================================
	fmt.Println("1. BREAK STATEMENT - Exit Loop Early")
	fmt.Println("-" + "---------------------------------")

	// Example 1: Break in for loop
	fmt.Println("\nSearch and Break (Find number 30):")
	numbers := []int{10, 20, 30, 40, 50}
	for index, num := range numbers {
		fmt.Printf("Checking index %d: %d\n", index, num)
		if num == 30 {
			fmt.Printf("Found 30 at index %d! Breaking...\n", index)
			break // Exit the loop immediately
		}
	}

	// Example 2: Break with conditions
	fmt.Println("\nBreak when value exceeds limit:")
	limit := 35
	for i := 0; i < 10; i++ {
		value := i * 10
		fmt.Printf("Iteration %d: Value = %d\n", i, value)
		if value > limit {
			fmt.Printf("Value exceeded limit (%d > %d). Breaking!\n", value, limit)
			break
		}
	}

	// Example 3: Break in nested loops (breaks inner loop only)
	fmt.Println("\nBreak in Nested Loop:")
	fmt.Println("(Regular break only exits the innermost loop)")
	for i := 1; i <= 3; i++ {
		fmt.Printf("Outer Loop: i = %d\n", i)
		for j := 1; j <= 5; j++ {
			fmt.Printf("  Inner Loop: j = %d\n", j)
			if j == 3 {
				fmt.Println("  Breaking inner loop...")
				break // Only breaks inner loop, outer continues
			}
		}
	}

	// ============================================
	// 2. CONTINUE STATEMENT
	// ============================================
	fmt.Println("\n2. CONTINUE STATEMENT - Skip to Next Iteration")
	fmt.Println("-" + "---------------------------------")

	// Example 1: Skip even numbers
	fmt.Println("\nSkip Even Numbers (continue):")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Printf("Skipping even number: %d\n", i)
			continue // Jump to next iteration
		}
		fmt.Printf("Processing odd number: %d\n", i)
	}

	// Example 2: Skip specific items from slice
	fmt.Println("\nSkip Certain Items:")
	fruits := []string{"Apple", "Banana", "Cherry", "Date", "Elderberry"}
	skipFruit := "Cherry"
	for _, fruit := range fruits {
		if fruit == skipFruit {
			fmt.Printf("Skipping %s\n", fruit)
			continue
		}
		fmt.Printf("Processing: %s\n", fruit)
	}

	// Example 3: Continue in nested loops
	fmt.Println("\nContinue in Nested Loop:")
	fmt.Println("(Regular continue only skips current inner iteration)")
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			if j == 2 {
				fmt.Printf("  Skipping i=%d, j=%d\n", i, j)
				continue // Only skips inner loop iteration
			}
			fmt.Printf("  Processing i=%d, j=%d\n", i, j)
		}
	}

	// ============================================
	// 3. PRACTICAL USE CASES
	// ============================================
	fmt.Println("\n3. PRACTICAL USE CASES")
	fmt.Println("-" + "---------------------------------")

	// Use Case 1: Find first match and process
	fmt.Println("\nFind First Student with Grade >= 90:")
	students := map[string]int{
		"Alice": 95,
		"Bob":   87,
		"Carol": 92,
		"David": 88,
	}

	found := false
	for name, grade := range students {
		if grade >= 90 {
			fmt.Printf("Found: %s with grade %d\n", name, grade)
			found = true
			break
		}
	}
	if !found {
		fmt.Println("No student found with grade >= 90")
	}

	// Use Case 2: Process valid items, skip invalid ones
	fmt.Println("\nProcess Valid Ages (20-65):")
	ages := []int{15, 25, 30, 70, 35, 18, 45}
	validCount := 0
	for _, age := range ages {
		if age < 20 || age > 65 {
			fmt.Printf("Skipping invalid age: %d\n", age)
			continue
		}
		validCount++
		fmt.Printf("Valid age: %d\n", age)
	}
	fmt.Printf("Total valid ages: %d\n", validCount)

	// Use Case 3: Early exit from loop
	fmt.Println("\nPassword Attempt (max 3 tries):")
	correctPassword := "GoLang123"
	attempts := 0
	maxAttempts := 3

	passwords := []string{"wrong1", "wrong2", "GoLang123", "wrong4"}

	for _, password := range passwords {
		attempts++
		fmt.Printf("Attempt %d: Trying '%s'...\n", attempts, password)

		if password == correctPassword {
			fmt.Println("✓ Password correct! Access granted.")
			break
		}

		if attempts >= maxAttempts {
			fmt.Printf("✗ Maximum attempts (%d) reached. Access denied.\n", maxAttempts)
			break
		}
	}

	// ============================================
	// 4. BREAK VS CONTINUE COMPARISON
	// ============================================
	fmt.Println("\n4. BREAK VS CONTINUE - SIDE BY SIDE")
	fmt.Println("-" + "---------------------------------")

	fmt.Println("\nUsing BREAK:")
	fmt.Println("Numbers: ", []int{1, 2, 3, 4, 5})
	for _, n := range []int{1, 2, 3, 4, 5} {
		if n == 3 {
			fmt.Println("Breaking at 3...")
			break
		}
		fmt.Printf("%d ", n)
	}
	fmt.Println("\n(Output: 1 2) - Loop exits completely\n")

	fmt.Println("Using CONTINUE:")
	fmt.Println("Numbers: ", []int{1, 2, 3, 4, 5})
	for _, n := range []int{1, 2, 3, 4, 5} {
		if n == 3 {
			fmt.Println("Skipping 3...")
			continue
		}
		fmt.Printf("%d ", n)
	}
	fmt.Println("\n(Output: 1 2 4 5) - Loop continues\n")

	// ============================================
	// 5. LABELED BREAK (Advanced)
	// ============================================
	fmt.Println("\n5. LABELED BREAK - Break Outer Loop (Advanced)")
	fmt.Println("-" + "---------------------------------")

	fmt.Println("\nWithout Label (breaks inner loop only):")
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Printf("i=%d, j=%d ", i, j)
			if i == 2 && j == 2 {
				fmt.Println("(break)")
				break // Only breaks inner loop
			}
		}
		fmt.Println()
	}

	fmt.Println("\nWith Label (breaks to labeled point):")
	OuterLoop2:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Printf("i=%d, j=%d ", i, j)
			if i == 2 && j == 2 {
				fmt.Println("(break OuterLoop2)")
				break OuterLoop2 // Breaks out of outer loop
			}
		}
		fmt.Println()
	}

	// ============================================
	// 6. COMMON PATTERNS
	// ============================================
	fmt.Println("\n6. COMMON PATTERNS")
	fmt.Println("-" + "---------------------------------")

	// Pattern 1: Validate and process
	fmt.Println("\nValidate Items Before Processing:")
	items := []struct {
		name  string
		price float64
	}{
		{"Laptop", 1200.50},
		{"Mouse", 25.00},
		{"Keyboard", -50.00}, // Invalid: negative price
		{"Monitor", 350.00},
	}

	for _, item := range items {
		// Skip invalid items
		if item.price <= 0 {
			fmt.Printf("Skipping invalid item %s (price: %.2f)\n", item.name, item.price)
			continue
		}
		fmt.Printf("Processing: %s - $%.2f\n", item.name, item.price)
	}

	// Pattern 2: Process until condition
	fmt.Println("\nProcess Until Running Sum Exceeds 100:")
	total := 0
	values := []int{20, 15, 35, 40, 50, 30}
	for _, val := range values {
		total += val
		fmt.Printf("Added %d, Total: %d\n", val, total)
		if total > 100 {
			fmt.Println("Total exceeded 100. Stopping.")
			break
		}
	}
}

