package main

import (
	"fmt"
)

// 20. Valid Parentheses
//
// Problem Statement:
// Given string s containing just characters '(', ')', '{', '}', '[', ']',
// determine if input string is valid. String valid if brackets closed by
// same type of brackets, in correct order.
//
// Example 1:
// Input:  s = "()"
// Output: true
//
// Example 2:
// Input:  s = "()[]{}"
// Output: true
//
// Example 3:
// Input:  s = "(]"
// Output: false
//
// Example 4:
// Input:  s = "([)]"
// Output: false
//
// Constraints:
// 1 <= s.length <= 10^4
// s consists of parentheses only '()[]{}'

func isValid(s string) bool {
	var stack []byte
	for i := range s {
		switch s[i] {
		case '{':
			stack = append(stack, '}')
		case '[':
			stack = append(stack, ']')
		case '(':
			stack = append(stack, ')')
		case '}', ']', ')':
			if len(stack) == 0 || stack[len(stack)-1] != s[i] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

func main() {
	fmt.Println(isValid("()"))
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("(]"))
	fmt.Println(isValid("([)]"))
	fmt.Println(isValid(")"))
}
