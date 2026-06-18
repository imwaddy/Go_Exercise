package main

import (
	"fmt"
)

func main() {
	myInt := "III"

	fmt.Println("My string: ", romanToInt(myInt))
}

var m = map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}

func romanToInt(s string) int {
	out := 0
	ln := len(s)
	for i := 0; i < ln; i++ {
		c := string(s[i])
		vc := m[c]
		if i < ln-1 {
			cnext := string(s[i+1])
			vcnext := m[cnext]
			if vc < vcnext {
				out += vcnext - vc
				i++
			} else {
				out += vc
			}
		} else {
			out += vc
		}
	}
	return out
}

// var numbers = map[rune]int{
// 	'I': 1,
// 	'V': 5,
// 	'X': 10,
// 	'L': 50,
// 	'C': 100,
// 	'D': 500,
// 	'M': 1000,
// }

// func romanToInt(s string) int {
// 	result := numbers[rune(s[len(s)-1])]
// 	for i := 0; i < len(s)-1; i++ {
// 		num := numbers[rune(s[i])]
// 		next := numbers[rune(s[i+1])]
// 		if num < next {
// 			result -= num
// 		} else {
// 			result += num
// 		}
// 	}
// 	return result
// }
