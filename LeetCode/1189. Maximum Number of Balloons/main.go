package main

import (
	"fmt"
)

// 1189. Maximum Number of Balloons
//
// Problem Statement:
// Given string text, use characters of text to form as many instances of
// word "balloon" as possible. Each character in text can be used at most
// once. Return max number of instances of "balloon" that can be formed.
//
// Example 1:
// Input:  text = "nlaebolko"
// Output: 1
//
// Example 2:
// Input:  text = "loonbalxballpoon"
// Output: 2
//
// Example 3:
// Input:  text = "leetcode"
// Output: 0
//
// Constraints:
// 1 <= text.length <= 10^4
// text consists of lowercase English letters

func maxNumberOfBalloons(text string) int {
	m := make(map[rune]int)
	for _, ch := range text {
		m[ch] += 1
	}

	needed := map[rune]int{'b': 1, 'a': 1, 'l': 2, 'o': 2, 'n': 1}

	min := -1
	for ch, need := range needed {
		ratio := m[ch] / need
		if min == -1 || ratio < min {
			min = ratio
		}
	}

	return min
}

func main() {
	fmt.Println(maxNumberOfBalloons("nlaebolko"))
}
