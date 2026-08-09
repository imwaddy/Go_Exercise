package main

import (
	"fmt"
)

// 268. Missing Number
//
// Problem Statement:
// Given array nums containing n distinct numbers in range [0, n], return
// the only number in range missing from array.
//
// Example 1:
// Input:  nums = [3,0,1]
// Output: 2
//
// Example 2:
// Input:  nums = [0,1]
// Output: 2
//
// Example 3:
// Input:  nums = [9,6,4,2,3,5,7,0,1]
// Output: 8
//
// Constraints:
// n == nums.length
// 1 <= n <= 10^4
// 0 <= nums[i] <= n
// All numbers in nums are unique

func missingNumber(nums []int) int {
	sum := len(nums)
	for i, num := range nums {
		sum ^= i ^ num
	}
	return sum
}

func main() {
	fmt.Println(missingNumber([]int{3, 0, 1}))
	fmt.Println(missingNumber([]int{0, 1}))
	fmt.Println(missingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}))
}
