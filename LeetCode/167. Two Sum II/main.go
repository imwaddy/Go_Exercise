package main

import (
	"fmt"
	"slices"
)

// 167. Two Sum II - Input Array Is Sorted
//
// Problem Statement:
// Given 1-indexed array numbers sorted in non-decreasing order, find two
// numbers that add up to target. Return their indices as [index1, index2],
// where 1 <= index1 < index2 <= numbers.length.
//
// Assume exactly one solution exists, and same element cannot be used twice.
// Must use only constant extra space.
//
// Example 1:
// Input:  numbers = [2,7,11,15], target = 9
// Output: [1,2]
// (numbers[1] + numbers[2] = 2 + 7 = 9)
//
// Example 2:
// Input:  numbers = [2,3,4], target = 6
// Output: [1,3]
//
// Example 3:
// Input:  numbers = [-1,0], target = -1
// Output: [1,2]
//
// Constraints:
// 2 <= numbers.length <= 3 * 10^4
// -1000 <= numbers[i] <= 1000
// numbers sorted in non-decreasing order
// -1000 <= target <= 1000
// exactly one valid answer exists
//
// Hint:
// Two-pointer approach. left at start, right at end. If sum too big, move
// right inward. If sum too small, move left inward. Stop when sum matches.

func TwoSum(numbers []int, target int) []int {
	var left = 0
	var right = len(numbers) - 1

	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			return []int{left + 1, right + 1}
		} else if sum > target {
			right--
		} else {
			left++
		}
	}
	return nil
}

func main() {
	tests := []struct {
		numbers []int
		target  int
		want    []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{1, 2}},
		{[]int{2, 3, 4}, 6, []int{1, 3}},
		{[]int{-1, 0}, -1, []int{1, 2}},
	}

	for _, tc := range tests {
		got := TwoSum(tc.numbers, tc.target)
		status := "PASS"
		if !slices.Equal(got, tc.want) {
			status = "FAIL"
		}
		fmt.Printf("[%s] got=%v want=%v numbers=%v target=%v\n", status, got, tc.want, tc.numbers, tc.target)
	}
}
