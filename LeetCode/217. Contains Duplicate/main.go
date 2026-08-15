package main

import (
	"fmt"
)

// 217. Contains Duplicate
//
// Problem Statement:
// Given integer array nums, return true if any value appears at least
// twice in the array, and false if every element is distinct.
//
// Example 1:
// Input:  nums = [1,2,3,1]
// Output: true
//
// Example 2:
// Input:  nums = [1,2,3,4]
// Output: false
//
// Example 3:
// Input:  nums = [1,1,1,3,3,4,3,2,4,2]
// Output: true
//
// Constraints:
// 1 <= nums.length <= 10^5
// -10^9 <= nums[i] <= 10^9

func containsDuplicate(nums []int) bool {
	m := make(map[int]struct{}, 0)
	for _, num := range nums {
		_, exists := m[num]
		if exists {
			return true
		}
		m[num] = struct{}{}
	}
	return false
}

func main() {
	fmt.Println(containsDuplicate([]int{1, 2, 3, 1}))
	fmt.Println(containsDuplicate([]int{1, 2, 3, 4}))
	fmt.Println(containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}))
}
