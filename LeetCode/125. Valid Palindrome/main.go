package main

import (
	"fmt"
	"strings"
	"unicode"
)

// 125. Valid Palindrome
//
// Problem Statement:
// Given string s, check if it is palindrome, considering only alphanumeric
// characters and ignoring case.
//
// Example 1:
// Input:  s = "A man, a plan, a canal: Panama"
// Output: true
//
// Example 2:
// Input:  s = "race a car"
// Output: false
//
// Example 3:
// Input:  s = " "
// Output: true
//
// Constraints:
// 1 <= s.length <= 2 * 10^5
// s consists only of printable ASCII characters
//
// Hint:
// Two-pointer approach. left from start, right from end. Skip non-alphanumeric
// chars on both sides, compare lowercased chars, move inward till pointers meet.

func sanitizeString(s string) string {
	var str strings.Builder
	for _, r := range s {
		if unicode.IsDigit(r) || unicode.IsLetter(r) {
			str.WriteRune(unicode.ToLower(r))
		}
	}
	return str.String()
}

func IsPalindrome(s string) bool {
	s = sanitizeString(s)
	if len(s) == 0 {
		return true
	}

	runes := []rune(s)
	var left = 0
	var right = len(runes) - 1

	for left < right {
		if runes[left] != runes[right] {
			return false
		}
		left++
		right--
	}

	return true
}

func main() {
	tests := []struct {
		input string
		want  bool
	}{
		{"A man, a plan, a canal: Panama", true},
		{"race a car", false},
		{" ", true},
	}

	for _, tc := range tests {
		got := IsPalindrome(tc.input)
		status := "PASS"
		if got != tc.want {
			status = "FAIL"
		}
		fmt.Printf("[%s] input=%q got=%v want=%v\n", status, tc.input, got, tc.want)
	}
}
