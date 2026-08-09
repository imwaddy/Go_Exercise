package main

import (
	"fmt"
	"slices"
)

// 1365. How Many Numbers Are Smaller Than the Current Number
//
// Problem Statement:
// Given array nums, for each nums[i] count how many numbers in array are
// smaller than it. Return array of counts.
//
// Example 1:
// Input:  nums = [8,1,2,2,3]
// Output: [4,0,1,1,3]
//
// Example 2:
// Input:  nums = [6,5,4,8]
// Output: [2,1,0,3]
//
// Example 3:
// Input:  nums = [7,7,7,7]
// Output: [0,0,0,0]
//
// Constraints:
// 2 <= nums.length <= 500
// 0 <= nums[i] <= 100

func smallerNumbersThanCurrent(nums []int) []int {
	result := make([]int, len(nums))
	newArray := make([]int, len(nums))

	copy(newArray, nums)

	slices.Sort(newArray)

	m := make(map[int]int)

	for i, num := range newArray {
		_, exists := m[num]
		if !exists {
			m[num] = i
		}
	}

	for i, v := range nums {
		val, _ := m[v]
		result[i] = val
	}

	return result
}

func main() {
	fmt.Println(smallerNumbersThanCurrent([]int{8, 1, 2, 2, 3}))
}
