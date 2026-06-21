package main

import "fmt"

var stack []string

func parenthesisCombination() {
	str := "abc[]["

	for _, s := range str {

		switch s {
		case '[', '{', '(':
			stack = append(stack, string(s))
			fmt.Println(stack)
		default:
			switch s {
			case '}':
				if stack[len(stack)-1] == "{" {
					stack = stack[:len(stack)-1]
				}
			case ']':
				if stack[len(stack)-1] == "[" {
					stack = stack[:len(stack)-1]
				}
			case ')':
				if stack[len(stack)-1] == "(" {
					stack = stack[:len(stack)-1]
				}
			}
		}

	}
	if len(stack) != 0 {
		fmt.Println("valid")
	} else {
		fmt.Println("invalid")
	}

}
