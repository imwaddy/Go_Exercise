package main

import "fmt"

// 9. Palindrome Number
//
// Problem Statement:
// Given integer x, return true if x is palindrome integer (reads same
// forward and backward).
//
// Example 1:
// Input:  x = 121
// Output: true
//
// Example 2:
// Input:  x = -121
// Output: false
// (reads -121- backward, negative sign breaks it)
//
// Example 3:
// Input:  x = 10
// Output: false
//
// Constraints:
// -2^31 <= x <= 2^31 - 1

func isPalindrome(x int) bool {
	var rev = 0
	var original = x
	for x != 0 {
		rem := x % 10
		rev = (rev * 10) + rem
		if rev < 0 {
			return false
		}
		x = x / 10
	}
	return rev == original
}

func main() {
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(-121))
	fmt.Println(isPalindrome(10))
}
