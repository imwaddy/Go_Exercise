package main

import (
	"fmt"
	"slices"
)

// 344. Reverse String
//
// Problem Statement:
// Write function that reverses a string. Input string given as array of
// characters s. Must modify input array in-place with O(1) extra memory.
//
// Example 1:
// Input:  s = ['h','e','l','l','o']
// Output: ['o','l','l','e','h']
//
// Example 2:
// Input:  s = ['H','a','n','n','a','h']
// Output: ['h','a','n','n','a','H']
//
// Constraints:
// 1 <= s.length <= 10^5
// s[i] is a printable ASCII character
//
// Hint:
// Two-pointer approach. left at start, right at end. Swap characters, move
// both pointers inward, stop when they meet or cross.

func ReverseString(s []byte) {
	first := 0
	last := len(s) - 1

	for first < last {
		s[first], s[last] = s[last], s[first]
		first++
		last--
	}
}

func main() {
	tests := []struct {
		input []byte
		want  []byte
	}{
		{[]byte{'h', 'e', 'l', 'l', 'o'}, []byte{'o', 'l', 'l', 'e', 'h'}},
		{[]byte{'H', 'a', 'n', 'n', 'a', 'h'}, []byte{'h', 'a', 'n', 'n', 'a', 'H'}},
	}

	for _, tc := range tests {
		ReverseString(tc.input)
		status := "PASS"
		if !slices.Equal(tc.input, tc.want) {
			status = "FAIL"
		}
		fmt.Printf("[%s] got=%q want=%q\n", status, string(tc.input), string(tc.want))
	}
}
