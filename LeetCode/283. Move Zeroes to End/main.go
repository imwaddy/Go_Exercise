package main

import "fmt"

// 283. Move Zeroes to End
//
// Problem Statement:
// Given integer array nums, move all zeroes to end of it, while preserving
// relative order of non-zero elements.
//
// Must do this in-place without making copy of array.
//
// Example 1:
// Input:  nums = [0,1,0,3,12]
// Output: [1,3,12,0,0]
//
// Example 2:
// Input:  nums = [0]
// Output: [0]
//
// Example 3:
// Input:  nums = [1,2,3]
// Output: [1,2,3]
//
// Constraints:
// 1 <= nums.length <= 10^4
// -2^31 <= nums[i] <= 2^31 - 1
//
// Hint:
// Two-pointer approach. Track position where next non-zero element belongs,
// swap as you scan array once.

func MoveZeroes(nums []int) {
	var left = 0
	var right = 0

	// var m = make(map[int]struct{}, 0)
	for right <= len(nums)-1 {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}

	fmt.Println(nums)
}

func main() {
	MoveZeroes([]int{0, 1, 0, 3, 12})
	MoveZeroes([]int{0, 1, 0, 3, 12})
	MoveZeroes([]int{0})
	MoveZeroes([]int{1, 2, 3})
}
