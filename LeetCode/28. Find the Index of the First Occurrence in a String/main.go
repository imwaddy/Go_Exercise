package main

import "fmt"

// 28. Find the Index of the First Occurrence in a String
//
// Problem Statement:
// Given two strings needle and haystack, return index of first occurrence
// of needle in haystack, or -1 if needle not part of haystack.
//
// Example 1:
// Input:  haystack = "sadbutsad", needle = "sad"
// Output: 0
//
// Example 2:
// Input:  haystack = "leetcode", needle = "leeto"
// Output: -1
//
// Constraints:
// 1 <= haystack.length, needle.length <= 10^4
// haystack and needle consist of only lowercase English characters

func strStr(haystack string, needle string) int {
	i := 0
	j := len(needle)
	for j <= len(haystack) {
		if haystack[i:j] == needle {
			return i
		}
		i++
		j++
	}
	return -1
}

func main() {
	fmt.Println(strStr("sadbutsad", "sad"))
	fmt.Println(strStr("leetcode", "leeto"))
	fmt.Println(strStr("a", "a"))
	fmt.Println(strStr("mississippi", "issip"))
	fmt.Println(strStr("aaaaa", "bba"))
	fmt.Println(strStr("hello", "ll"))
}
