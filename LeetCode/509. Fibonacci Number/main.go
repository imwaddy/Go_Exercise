package main

import "fmt"

func check(name string, got int, want int) {
	if got == want {
		fmt.Printf("%s PASS got=%v\n", name, got)
	} else {
		fmt.Printf("%s FAIL got=%v want=%v\n", name, got, want)
	}
}

// 509. Fibonacci Number
//
// Problem Statement:
// The Fibonacci numbers form a sequence: F(0) = 0, F(1) = 1,
// F(n) = F(n-1) + F(n-2) for n > 1. Given n, return F(n).
//
// Example 1:
// Input:  n = 2
// Output: 1
//
// Example 2:
// Input:  n = 3
// Output: 2
//
// Example 3:
// Input:  n = 4
// Output: 3
//
// Constraints:
// 0 <= n <= 30

func fib(n int) int {
	if n <= 1 {
		return n
	}

	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

func main() {
	check("[1]", fib(2), 1)
	check("[2]", fib(3), 2)
	check("[3]", fib(4), 3)
	check("[4]", fib(0), 0)
	check("[5]", fib(1), 1)
	check("[6]", fib(10), 55)
}
