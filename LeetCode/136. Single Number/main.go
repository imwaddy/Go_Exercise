package main

import (
	"fmt"
)

// 136. Single Number
//
// Problem Statement:
// Given non-empty array of integers nums, every element appears twice
// except for one. Find that single one.
//
// Example 1:
// Input:  nums = [2,2,1]
// Output: 1
//
// Example 2:
// Input:  nums = [4,1,2,1,2]
// Output: 4
//
// Example 3:
// Input:  nums = [1]
// Output: 1
//
// Constraints:
// 1 <= nums.length <= 3 * 10^4
// -3 * 10^4 <= nums[i] <= 3 * 10^4
// Each element appears twice except for one which appears once

func singleNumber(nums []int) int {
	var ans = 0
	for _, num := range nums {
		ans ^= num
	}
	return ans
}

func main() {
	fmt.Println(singleNumber([]int{2, 2, 1}))
	fmt.Println(singleNumber([]int{4, 1, 2, 1, 2}))
	fmt.Println(singleNumber([]int{1}))
}
