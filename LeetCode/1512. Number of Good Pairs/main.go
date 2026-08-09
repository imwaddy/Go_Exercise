package main

import (
	"fmt"
)

// 1512. Number of Good Pairs
//
// Problem Statement:
// Given array of integers nums, a pair (i, j) is good if nums[i] == nums[j]
// and i < j. Return number of good pairs.
//
// Example 1:
// Input:  nums = [1,2,3,1,1,3]
// Output: 4
//
// Example 2:
// Input:  nums = [1,1,1,1]
// Output: 6
//
// Example 3:
// Input:  nums = [1,2,3]
// Output: 0
//
// Constraints:
// 1 <= nums.length <= 100
// 1 <= nums[i] <= 100

func numIdenticalPairs(nums []int) int {
	m := map[int]int{}
	for _, num := range nums {
		m[num] += 1
	}
	var count int
	for _, v := range m {
		count += v * (v - 1) / 2
	}
	return count
}

func main() {
	fmt.Println(numIdenticalPairs([]int{1, 2, 3, 1, 1, 3}))
}
