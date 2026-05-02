package main

import (
	"fmt"
)

func main() {
	fmt.Println("=== LOGICAL OPERATORS IN GO ===\n")

	// ============================================
	// 1. AND OPERATOR (&&)
	// ============================================
	fmt.Println("1. AND OPERATOR (&&) - ALL conditions must be TRUE")
	fmt.Println("-" + "---------------------------------")

	age := 25
	income := 50000

	fmt.Println("\nExample: Check if person can get loan")
	fmt.Printf("Age: %d, Income: $%d\n", age, income)

	// AND operator - both conditions must be true
	if age >= 21 && income >= 30000 {
		fmt.Println("✓ Eligible for loan (age >= 21 AND income >= 30000)")
	} else {
		fmt.Println("✗ Not eligible for loan")
	}

	// Truth table for AND
	fmt.Println("\nAND Truth Table:")
	conditions := []struct {
		name   string
		a      bool
		b      bool
		result bool
	}{
		{"true && true", true, true, true},
		{"true && false", true, false, false},
		{"false && true", false, true, false},
		{"false && false", false, false, false},
	}

	for _, cond := range conditions {
		result := cond.a && cond.b
		fmt.Printf("%s = %v\n", cond.name, result)
	}

	// Practical example: Multiple conditions
	fmt.Println("\nMultiple AND Conditions:")
	score := 85
	attendance := 95
	behavior := 9

	if score >= 80 && attendance >= 90 && behavior >= 8 {
		fmt.Println("✓ Student is eligible for honors program")
	} else {
		fmt.Println("✗ Student does not meet all requirements")
	}

	// ============================================
	// 2. OR OPERATOR (||)
	// ============================================
	fmt.Println("\n2. OR OPERATOR (||) - AT LEAST ONE condition must be TRUE")
	fmt.Println("-" + "---------------------------------")

	day := "Saturday"
	isHoliday := false

	fmt.Println("\nExample: Is it a day off?")
	fmt.Printf("Day: %s, Holiday: %v\n", day, isHoliday)

	if day == "Saturday" || day == "Sunday" || isHoliday {
		fmt.Println("✓ It's a day off!")
	} else {
		fmt.Println("✗ It's a working day")
	}

	// Truth table for OR
	fmt.Println("\nOR Truth Table:")
	orConditions := []struct {
		name   string
		a      bool
		b      bool
		result bool
	}{
		{"true || true", true, true, true},
		{"true || false", true, false, true},
		{"false || true", false, true, true},
		{"false || false", false, false, false},
	}

	for _, cond := range orConditions {
		result := cond.a || cond.b
		fmt.Printf("%s = %v\n", cond.name, result)
	}

	// Practical example: Multiple conditions
	fmt.Println("\nMultiple OR Conditions:")
	userRole := "admin"
	isOwner := false
	isModerator := false

	if userRole == "admin" || isOwner || isModerator {
		fmt.Println("✓ User has elevated privileges")
	} else {
		fmt.Println("✗ User has standard privileges")
	}

	// ============================================
	// 3. NOT OPERATOR (!)
	// ============================================
	fmt.Println("\n3. NOT OPERATOR (!) - NEGATES a condition")
	fmt.Println("-" + "---------------------------------")

	isRaining := true
	canGoOutside := !isRaining // NOT operator

	fmt.Println("\nExample: Can I go outside?")
	fmt.Printf("Is it raining: %v\n", isRaining)
	fmt.Printf("Can go outside: %v\n", canGoOutside)

	if !isRaining {
		fmt.Println("✓ You can go outside! It's not raining.")
	} else {
		fmt.Println("✗ You cannot go outside. It's raining.")
	}

	// NOT with variables
	fmt.Println("\nNOT Examples:")
	isLoggedIn := false
	if !isLoggedIn {
		fmt.Println("✓ Please log in")
	}

	hasPermission := false
	if !hasPermission {
		fmt.Println("✓ Access denied. You don't have permission.")
	}

	// ============================================
	// 4. COMBINING OPERATORS (Precedence)
	// ============================================
	fmt.Println("\n4. COMBINING OPERATORS - Order Matters")
	fmt.Println("-" + "---------------------------------")

	fmt.Println("\nOperator Precedence (highest to lowest):")
	fmt.Println("1. ! (NOT)")
	fmt.Println("2. && (AND)")
	fmt.Println("3. || (OR)")

	fmt.Println("\nExample: Determine membership tier")
	totalSpending := 5000
	yearsCustomer := 3
	isVIP := false

	// Conditions:
	// - VIP if premium customer
	// - OR has spent > $3000 AND been customer for 2+ years
	if isVIP || (totalSpending > 3000 && yearsCustomer >= 2) {
		fmt.Println("✓ Eligible for Gold membership")
	} else {
		fmt.Println("✗ Not eligible for Gold membership")
	}

	// Show the evaluation step by step
	fmt.Println("\nStep-by-step evaluation:")
	fmt.Printf("isVIP = %v\n", isVIP)
	fmt.Printf("totalSpending > 3000 = %v\n", totalSpending > 3000)
	fmt.Printf("yearsCustomer >= 2 = %v\n", yearsCustomer >= 2)
	fmt.Printf("(totalSpending > 3000 && yearsCustomer >= 2) = %v\n", (totalSpending > 3000 && yearsCustomer >= 2))
	fmt.Printf("isVIP || (totalSpending > 3000 && yearsCustomer >= 2) = %v\n", isVIP || (totalSpending > 3000 && yearsCustomer >= 2))

	// ============================================
	// 5. SHORT-CIRCUIT EVALUATION
	// ============================================
	fmt.Println("\n5. SHORT-CIRCUIT EVALUATION - Optimization")
	fmt.Println("-" + "---------------------------------")

	fmt.Println("\nAND Short-circuit (stops if first is false):")
	value1 := 5
	value2 := 10

	if value1 > 10 && value2 > 5 {
		fmt.Println("Both conditions are true")
	} else {
		fmt.Println("First condition was false, so second wasn't evaluated")
	}

	fmt.Println("\nOR Short-circuit (stops if first is true):")
	if value1 == 5 || value2 == 100 {
		fmt.Println("First condition was true, so second wasn't evaluated")
	}

	// ============================================
	// 6. PRACTICAL EXAMPLES
	// ============================================
	fmt.Println("\n6. PRACTICAL EXAMPLES")
	fmt.Println("-" + "---------------------------------")

	// Example 1: Game level access
	fmt.Println("\nGame Level Access:")
	userLevel := 15
	hasMapUnlocked := true
	isPremium := false

	if (userLevel >= 10 || isPremium) && hasMapUnlocked {
		fmt.Println("✓ You can play this level!")
	} else {
		fmt.Println("✗ You don't have access to this level")
	}

	// Example 2: Password validation
	fmt.Println("\nPassword Validation:")
	password := "SecurePass123!"
	hasUppercase := len(password) > 0 && password[0] >= 'A' && password[0] <= 'Z'
	hasNumber := false
	hasSpecial := false

	// Check for number
	for _, char := range password {
		if char >= '0' && char <= '9' {
			hasNumber = true
		}
		if char == '!' || char == '@' || char == '#' {
			hasSpecial = true
		}
	}

	fmt.Printf("Password: %s\n", password)
	fmt.Printf("Has uppercase: %v\n", hasUppercase)
	fmt.Printf("Has number: %v\n", hasNumber)
	fmt.Printf("Has special char: %v\n", hasSpecial)

	if len(password) >= 8 && hasNumber && hasSpecial {
		fmt.Println("✓ Strong password")
	} else {
		fmt.Println("✗ Weak password")
	}

	// Example 3: Discount eligibility
	fmt.Println("\nDiscount Eligibility:")
	cartTotal := 150
	hasPromoCode := true
	isStudentVerified := false

	if (cartTotal >= 100 || hasPromoCode) && !isStudentVerified {
		discount := 10
		fmt.Printf("✓ You get %d%% discount!\n", discount)
	} else if isStudentVerified && hasPromoCode {
		discount := 20
		fmt.Printf("✓ You get %d%% student + promo discount!\n", discount)
	} else {
		fmt.Println("✗ No discount available")
	}

	// ============================================
	// 7. COMMON MISTAKES
	// ============================================
	fmt.Println("\n7. COMMON MISTAKES")
	fmt.Println("-" + "---------------------------------")

	fmt.Println("\n❌ WRONG - Using = instead of ==:")
	fmt.Println("   if age = 25 { } // ERROR!")

	fmt.Println("\n✓ CORRECT:")
	if age == 25 {
		fmt.Println("   age equals 25")
	}

	fmt.Println("\n❌ WRONG - Using 'and' instead of '&&':")
	fmt.Println("   if score >= 80 and attendance >= 90 { } // Go doesn't have 'and'")

	fmt.Println("\n✓ CORRECT:")
	if score >= 80 && attendance >= 90 {
		fmt.Println("   Both conditions met")
	}

	fmt.Println("\n❌ WRONG - Forgetting parentheses with mixed operators:")
	fmt.Println("   if a || b && c // ambiguous!")
	fmt.Println("   // Is it: (a || b) && c  OR  a || (b && c) ?")

	fmt.Println("\n✓ CORRECT:")
	a, b, c := true, false, true
	if (a || b) && c {
		fmt.Println("   Clear intention with parentheses")
	}

	// ============================================
	// 8. DE MORGAN'S LAWS (Advanced)
	// ============================================
	fmt.Println("\n8. DE MORGAN'S LAWS - Logical Equivalences")
	fmt.Println("-" + "---------------------------------")

	x := true
	y := false

	fmt.Println("\nDe Morgan's Law 1: !(a && b) == !a || !b")
	fmt.Printf("!(true && false) = %v\n", !(x && y))
	fmt.Printf("!true || !false = %v\n", !x || !y)
	fmt.Println("Both expressions give the same result!")

	fmt.Println("\nDe Morgan's Law 2: !(a || b) == !a && !b")
	fmt.Printf("!(true || false) = %v\n", !(x || y))
	fmt.Printf("!true && !false = %v\n", !x && !y)
	fmt.Println("Both expressions give the same result!")
}

