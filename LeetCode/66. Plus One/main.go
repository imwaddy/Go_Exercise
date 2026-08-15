package main

import "fmt"

// 66. Plus One
//
// Problem Statement:
// Given array of digits representing a non-negative integer (most
// significant digit first), increment integer by one and return resulting
// array of digits.
//
// Example 1:
// Input:  digits = [1,2,3]
// Output: [1,2,4]
//
// Example 2:
// Input:  digits = [4,3,2,1]
// Output: [4,3,2,2]
//
// Example 3:
// Input:  digits = [9]
// Output: [1,0]
//
// Constraints:
// 1 <= digits.length <= 100
// 0 <= digits[i] <= 9

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}

		digits[i] = 0
	}

	// Example: [9, 9] becomes [1, 0, 0]
	return append([]int{1}, digits...)
}

func main() {
	fmt.Println(plusOne([]int{1, 2, 3}))
	fmt.Println(plusOne([]int{4, 3, 2, 1}))
	fmt.Println(plusOne([]int{9}))
	fmt.Println(plusOne([]int{9, 9, 9}))
	fmt.Println(plusOne([]int{0}))
	fmt.Println(plusOne([]int{1, 9, 9}))
}
