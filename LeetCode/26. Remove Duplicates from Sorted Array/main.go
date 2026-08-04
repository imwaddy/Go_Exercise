package main

import (
	"fmt"
	"slices"
)

// 26. Remove Duplicates from Sorted Array
//
// Problem Statement:
// Given integer array nums sorted in non-decreasing order, remove duplicates
// in-place such that each unique element appears only once. Relative order
// of elements should stay same.
//
// Return k, number of unique elements. First k elements of nums should hold
// final result.
//
// Example 1:
// Input:  nums = [1,1,2]
// Output: k = 2, nums = [1,2,_]
//
// Example 2:
// Input:  nums = [0,0,1,1,1,2,2,3,3,4]
// Output: k = 5, nums = [0,1,2,3,4,_,_,_,_,_]
//
// Constraints:
// 1 <= nums.length <= 3 * 10^4
// -100 <= nums[i] <= 100
// nums sorted in non-decreasing order
//
// Hint:
// Two-pointer approach. slow pointer tracks last unique element position,
// fast pointer scans ahead. When fast finds new value, place it after slow.

func RemoveDuplicates(nums []int) int {
	var left = 1

	for right := 1; right < len(nums); right++ {
		if nums[right] != nums[left-1] {
			nums[left] = nums[right]
			left++
		}
	}

	return left
}

func main() {
	tests := []struct {
		input   []int
		want    int
		wantArr []int
	}{
		{[]int{1, 1, 2}, 2, []int{1, 2}},
		{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5, []int{0, 1, 2, 3, 4}},
	}

	for _, tc := range tests {
		got := RemoveDuplicates(tc.input)
		gotArr := tc.input[:got]

		status := "PASS"
		if got != tc.want || !slices.Equal(gotArr, tc.wantArr) {
			status = "FAIL"
		}
		fmt.Printf("[%s] got=%v want=%v gotArr=%v wantArr=%v\n", status, got, tc.want, gotArr, tc.wantArr)
	}
}
