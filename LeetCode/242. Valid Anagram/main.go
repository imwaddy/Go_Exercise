package main

import (
	"fmt"
)

// 242. Valid Anagram
//
// Problem Statement:
// Given two strings s and t, return true if t is an anagram of s,
// and false otherwise.
//
// Example 1:
// Input:  s = "anagram", t = "nagaram"
// Output: true
//
// Example 2:
// Input:  s = "rat", t = "car"
// Output: false
//
// Example 3:
// Input:  s = "a", t = "ab"
// Output: false
//
// Constraints:
// 1 <= s.length, t.length <= 5 * 10^4
// s and t consist of lowercase English letters

func isAnagram(s string, t string) bool {
	m := make(map[rune]int, 0)

	if len(s) != len(t) {
		return false
	}

	for _, ch := range s {
		m[ch] += 1
	}

	for _, ch := range t {
		m[ch] -= 1
	}

	for _, v := range m {
		if v != 0 {
			return false
		}
	}

	return true
}

// optimized
// func isAnagram(s string, t string) bool {
// 	if len(s) != len(t) {
//                 return false
//         }
//         var count [26]int
//         for i := 0; i < len(s); i++ {
//                 count[s[i]-'a']++
//                 count[t[i]-'a']--
//         }
//         for _, v := range count {
//                 if v != 0 {
//                         return false
//                 }
//         }
//         return true
// }

func main() {
	fmt.Println(isAnagram("anagram", "nagaram"))
	fmt.Println(isAnagram("rat", "car"))
	fmt.Println(isAnagram("a", "ab"))
	fmt.Println(isAnagram("aacc", "ccac")) // false
}
