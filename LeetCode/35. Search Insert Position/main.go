package main

import "fmt"

// 35. Search Insert Position
//
// Problem Statement:
// Given sorted array of distinct integers nums and target value, return
// index if target found. If not, return index where it would be if
// inserted in order.
//
// Example 1:
// Input:  nums = [1,3,5,6], target = 5
// Output: 2
//
// Example 2:
// Input:  nums = [1,3,5,6], target = 2
// Output: 1
//
// Example 3:
// Input:  nums = [1,3,5,6], target = 7
// Output: 4
//
// Constraints:
// 1 <= nums.length <= 10^4
// -10^4 <= nums[i] <= 10^4
// nums sorted in ascending order, distinct values
// -10^4 <= target <= 10^4

func searchInsert(nums []int, target int) int {
	if target <= nums[0] {
		return 0
	}

	for i := 1; i < len(nums); i++ {
		if nums[i-1] <= target && nums[i] >= target {
			return i
		}
	}
	return len(nums)
}

func main() {
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 5))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 2))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 7))
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 0))
	fmt.Println(searchInsert([]int{1}, 1))
	fmt.Println(searchInsert([]int{1}, 0))
	fmt.Println(searchInsert([]int{1}, 2))
}
