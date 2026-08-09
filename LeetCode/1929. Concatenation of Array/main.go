package main

import (
	"fmt"
)

// 1929. Concatenation of Array
//
// Problem Statement:
// Given integer array nums of length n, return array ans of length 2n where
// ans[i] == nums[i] and ans[i + n] == nums[i] for 0 <= i < n. (i.e. ans is
// nums concatenated with itself)
//
// Example 1:
// Input:  nums = [1,2,1]
// Output: [1,2,1,1,2,1]
//
// Example 2:
// Input:  nums = [1,3,2,1]
// Output: [1,3,2,1,1,3,2,1]
//
// Constraints:
// n == nums.length
// 1 <= n <= 1000
// 1 <= nums[i] <= 1000

func getConcatenation(nums []int) []int {
	last := len(nums)

	result := make([]int, len(nums)*2)

	for first, num := range nums {
		result[first] = num
		result[last] = num
		last++
	}

	return result
}

func main() {
	fmt.Println(getConcatenation([]int{1, 2, 3}))
}
