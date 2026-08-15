package main

import (
	"fmt"
	"math"
)

// 7. Reverse Integer
//
// Problem Statement:
// Given 32-bit signed integer x, return x with digits reversed. If
// reversing overflows 32-bit signed integer range, return 0.
//
// Example 1:
// Input:  x = 123
// Output: 321
//
// Example 2:
// Input:  x = -123
// Output: -321
//
// Example 3:
// Input:  x = 120
// Output: 21
//
// Example 4:
// Input:  x = 1534236469
// Output: 0
//
// Constraints:
// -2^31 <= x <= 2^31 - 1

func reverse(x int) int {
	var no int
	for x != 0 {
		rem := x % 10
		no = (no * 10) + rem
		if no > math.MaxInt32 || no < math.MinInt32 {
			return 0
		}
		x = x / 10
	}
	return no
}

func main() {
	fmt.Println(reverse(123))
	fmt.Println(reverse(-123))
	fmt.Println(reverse(120))
	fmt.Println(reverse(1534236469))
}
