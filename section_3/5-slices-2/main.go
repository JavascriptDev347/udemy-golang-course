package main

import (
	"fmt"
	"slices"
)

func main() {
	// advanced slices
	original := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("Original %v, Len %d, cap %d\n", original, len(original), cap(original))

	s1 := original[:5]
	fmt.Printf("S1 %v, Len %d, cap %d\n", s1, len(s1), cap(s1))

	s2 := original[5:7]
	fmt.Printf("S2 %v, Len %d, cap %d\n", s2, len(s2), cap(s2))

	s3 := original[7:]
	fmt.Printf("S3 %v, Len %d, cap %d\n", s3, len(s3), cap(s3))

	s4 := original[:]
	fmt.Printf("S4 %v, Len %d, cap %d\n", s4, len(s4), cap(s4))

	slices.Contains(s4, 8)
	slices.Insert(s4, 5, 42)
	fmt.Printf("Modified S4 %v, Len %d, cap %d\n", s4, len(s4), cap(s4))

}
