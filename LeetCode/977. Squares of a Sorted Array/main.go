package main

import (
	"fmt"
	"slices"
)

// 977. Squares of a Sorted Array
//
// Problem Statement:
// Given integer array nums sorted in non-decreasing order, return array of
// squares of each number, also sorted in non-decreasing order.
//
// Example 1:
// Input:  nums = [-4,-1,0,3,10]
// Output: [0,1,9,16,100]
//
// Example 2:
// Input:  nums = [-7,-3,2,3,11]
// Output: [4,9,9,49,121]
//
// Constraints:
// 1 <= nums.length <= 10^4
// -10^4 <= nums[i] <= 10^4
// nums sorted in non-decreasing order
//
// Hint:
// Two-pointer approach. left at start, right at end. Largest square always
// comes from either end (most negative or most positive). Fill result array
// from back to front, moving whichever pointer gives bigger square inward.

func SortedSquares(nums []int) []int {
	n := len(nums)
	result := make([]int, n)
	left, right := 0, n-1
	for i := n - 1; i >= 0; i-- {
		if nums[left]*nums[left] > nums[right]*nums[right] {
			result[i] = nums[left] * nums[left]
			left++
		} else {
			result[i] = nums[right] * nums[right]
			right--
		}
	}
	return result
}

func main() {
	tests := []struct {
		input []int
		want  []int
	}{
		{[]int{-4, -1, 0, 3, 10}, []int{0, 1, 9, 16, 100}},
		{[]int{-7, -3, 2, 3, 11}, []int{4, 9, 9, 49, 121}},
	}

	for _, tc := range tests {
		got := SortedSquares(tc.input)
		status := "PASS"
		if !slices.Equal(got, tc.want) {
			status = "FAIL"
		}
		fmt.Printf("[%s] got=%v want=%v input=%v\n", status, got, tc.want, tc.input)
	}
}
