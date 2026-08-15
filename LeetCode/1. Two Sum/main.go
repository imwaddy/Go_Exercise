package main

import (
	"fmt"
)

// 1. Two Sum
//
// Problem Statement:
// Given array of integers nums and an integer target, return indices of
// the two numbers such that they add up to target. Each input has exactly
// one solution, and you may not use the same element twice.
//
// Example 1:
// Input:  nums = [2,7,11,15], target = 9
// Output: [0,1]
//
// Example 2:
// Input:  nums = [3,2,4], target = 6
// Output: [1,2]
//
// Example 3:
// Input:  nums = [3,3], target = 6
// Output: [0,1]
//
// Constraints:
// 2 <= nums.length <= 10^4
// -10^9 <= nums[i], target <= 10^9
// Exactly one valid answer exists

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, 0)
	for i, num := range nums {
		k := target - num
		val, exists := m[k]
		if exists {
			return []int{val, i}
		}
		m[num] = i

	}
	return nil
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
	fmt.Println(twoSum([]int{3, 3}, 6))
}
