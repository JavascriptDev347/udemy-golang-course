package main

import "fmt"

func twoSum(nums []int, target int) []int {

	length := len(nums)

	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func main() {
	nums := [4]int{2, 7, 11, 15}
	target := 9

	result := twoSum(nums[:], target)
	fmt.Println("Result: ", result)
}
