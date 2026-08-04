package main

import (
	"fmt"
	"strconv"
)

// 415. Add Strings
//
// Problem Statement:
// Given two non-negative integers num1 and num2 represented as strings,
// return sum of num1 and num2 as a string.
//
// Must not use any built-in library for handling big integers (no
// converting entire string to int type directly), and must not convert
// inputs to integer directly.
//
// Example 1:
// Input:  num1 = "11", num2 = "123"
// Output: "134"
//
// Example 2:
// Input:  num1 = "456", num2 = "77"
// Output: "533"
//
// Example 3:
// Input:  num1 = "0", num2 = "0"
// Output: "0"
//
// Constraints:
// 1 <= num1.length, num2.length <= 10^4
// num1 and num2 consist only of digits
// num1 and num2 don't have leading zeros, except for zero itself
//
// Hint:
// Two-pointer approach. Start from end of both strings (least significant
// digit), add digit-by-digit like manual addition, track carry, move
// pointers left until both exhausted.

func AddStrings(num1 string, num2 string) string {
	i := len(num1) - 1
	j := len(num2) - 1
	carry := 0
	result := ""

	for i >= 0 || j >= 0 || carry > 0 {
		d1, d2 := 0, 0
		if i >= 0 {
			d1 = int(num1[i] - '0')
		}
		if j >= 0 {
			d2 = int(num2[j] - '0')
		}

		sum := d1 + d2 + carry
		carry = sum / 10
		digit := sum % 10

		result = strconv.Itoa(digit) + result

		i--
		j--
	}

	return result
}

func main() {
	tests := []struct {
		num1 string
		num2 string
		want string
	}{
		{"11", "123", "134"},
		{"456", "77", "533"},
		{"0", "0", "0"},
	}

	for _, tc := range tests {
		got := AddStrings(tc.num1, tc.num2)
		status := "PASS"
		if got != tc.want {
			status = "FAIL"
		}
		fmt.Printf("[%s] got=%q want=%q num1=%q num2=%q\n", status, got, tc.want, tc.num1, tc.num2)
	}
}
