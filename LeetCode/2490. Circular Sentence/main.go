package main

import (
	"fmt"
	"strings"
)

// 2490. Circular Sentence
//
// Problem Statement:
// Sentence is list of words separated by single space, no leading/trailing
// spaces. Each word only has lowercase/uppercase letters, no punctuation.
//
// Sentence is circular if last char of a word equals first char of next
// word, including last word's last char equals first word's first char.
//
// Given string sentence, return true if it is circular, else false.
//
// Example 1:
// Input:  sentence = "leetcode exercises sound delightful"
// Output: true
// (leetcode[-1]='e', exercises[0]='e'; exercises[-1]='s', sound[0]='s';
// sound[-1]='d', delightful[0]='d'; delightful[-1]='l', leetcode[0]='l')
//
// Example 2:
// Input:  sentence = "eetcode"
// Output: true
// (single word, last char 'e' equals first char 'e')
//
// Example 3:
// Input:  sentence = "Leetcode is cool"
// Output: false
// (Leetcode[-1]='e', is[0]='i' -> mismatch)
//
// Constraints:
// 1 <= sentence.length <= 500
// sentence consists of only lowercase and uppercase English letters and spaces
// words in sentence separated by single space
// sentence has no leading/trailing spaces

func IsCircularSentence(sentence string) bool {
	words := strings.Split(sentence, " ")

	for i, word := range words {
		next := words[(i+1)%len(words)]
		if word[len(word)-1] != next[0] {
			return false
		}
	}
	return true
}

func main() {
	tests := []struct {
		input string
		want  bool
	}{
		{"leetcode exercises sound delightful", true},
		{"eetcode", true},
		{"Leetcode is cool", false},
	}

	for _, tc := range tests {
		got := IsCircularSentence(tc.input)
		status := "PASS"
		if got != tc.want {
			status = "FAIL"
		}
		fmt.Printf("[%s] got=%v want=%v input=%q\n", status, got, tc.want, tc.input)
	}
}
