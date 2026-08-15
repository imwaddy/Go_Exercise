package main

import (
	"fmt"
	"strings"
)

// 58. Length of Last Word
//
// Problem Statement:
// Given a string s consisting of words and spaces, return length of the
// last word in the string. A word is a maximal substring consisting of
// non-space characters only.
//
// Example 1:
// Input:  s = "Hello World"
// Output: 5
//
// Example 2:
// Input:  s = "   fly me   to   the moon  "
// Output: 4
//
// Example 3:
// Input:  s = "luffy is still joyboy"
// Output: 6
//
// Constraints:
// 1 <= s.length <= 10^4
// s consists of only English letters and spaces ' '
// there will be at least one word in s

func lengthOfLastWord(s string) int {
	arr := strings.Fields(s)
	return len(arr[len(arr)-1])
}

func main() {
	fmt.Println(lengthOfLastWord("Hello World"))
	fmt.Println(lengthOfLastWord("   fly me   to   the moon  "))
	fmt.Println(lengthOfLastWord("luffy is still joyboy"))
	fmt.Println(lengthOfLastWord("a"))
	fmt.Println(lengthOfLastWord("  a  "))
}
