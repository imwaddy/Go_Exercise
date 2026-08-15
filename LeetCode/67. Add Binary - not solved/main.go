package main

import (
	"fmt"
)

// 67. Add Binary
//
// Problem Statement:
// Given two binary strings a and b, return their sum as a binary string.
//
// Example 1:
// Input:  a = "11", b = "1"
// Output: "100"
//
// Example 2:
// Input:  a = "1010", b = "1011"
// Output: "10101"
//
// Constraints:
// 1 <= a.length, b.length <= 10^4
// a and b consist only of '0' or '1' characters
// each string does not contain leading zeros except for the zero itself
func addBinary(a string, b string) string {
	return ""
}

func main() {
	fmt.Println(addBinary("11", "1"))
	fmt.Println(addBinary("1010", "1011"))
	fmt.Println(addBinary("0", "0"))
	fmt.Println(addBinary("1", "111"))
}
