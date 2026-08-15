package main

import "fmt"

// 14. Longest Common Prefix
//
// Problem Statement:
// Given array of strings strs, return longest common prefix string among
// all strings. If no common prefix exists, return empty string "".
//
// Example 1:
// Input:  strs = ["flower","flow","flight"]
// Output: "fl"
//
// Example 2:
// Input:  strs = ["dog","racecar","car"]
// Output: ""
//
// Constraints:
// 1 <= strs.length <= 200
// 0 <= strs[i].length <= 200
// strs[i] consists of only lowercase English letters

func commonPrefix(common, str string) string {
	minLen := len(common)
	if minLen > len(str) {
		minLen = len(str)
	}
	i := 0
	for i < minLen && common[i] == str[i] {
		i++
	}
	return str[:i]
}

func longestCommonPrefix(strs []string) string {
	var common = strs[0]
	for i := 1; i < len(strs); i++ {
		common = commonPrefix(common, strs[i])
		if common == "" {
			return ""
		}
	}
	return common
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))
}
