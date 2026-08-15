package main

import "fmt"

var stack []string

func parenthesisCombination() {
	var str string = "abc[]["

	var stack []byte
	for i := range str {
		switch str[i] {
		case '{':
			stack = append(stack, '}')
		case '[':
			stack = append(stack, ']')
		case '(':
			stack = append(stack, ')')
		case '}', ']', ')':
			if len(stack) == 0 || stack[len(stack)-1] != str[i] {
				fmt.Println("Invalid")
			}
			stack = stack[:len(stack)-1]
		}
	}
	fmt.Println("Valid")
}
