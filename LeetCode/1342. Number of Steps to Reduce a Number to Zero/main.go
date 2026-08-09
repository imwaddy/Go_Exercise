package main

import (
	"fmt"
)

// 1342. Number of Steps to Reduce a Number to Zero
//
// Problem Statement:
// Given non-negative integer num, return number of steps to reduce it to
// zero. In one step: if num even, divide by 2; if odd, subtract 1.
//
// Example 1:
// Input:  num = 14
// Output: 6
//
// Example 2:
// Input:  num = 8
// Output: 4
//
// Example 3:
// Input:  num = 123
// Output: 12
//
// Constraints:
// 0 <= num <= 10^6

func numberOfSteps(num int) int {
	var count int
	for num > 0 {
		if num%2 == 0 {
			count++
			num = num / 2
		} else {
			num -= 1
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(numberOfSteps(14))
}
