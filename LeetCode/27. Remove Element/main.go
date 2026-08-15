package main

import "fmt"

// 27. Remove Element
//
// Problem Statement:
// Given array nums and value val, remove all occurrences of val in nums
// in-place. Return k, number of elements not equal to val. Elements beyond
// k don't matter.
//
// Example 1:
// Input:  nums = [3,2,2,3], val = 3
// Output: k = 2, nums = [2,2,_,_]
//
// Example 2:
// Input:  nums = [0,1,2,2,3,0,4,2], val = 2
// Output: k = 5, nums = [0,1,4,0,3,_,_,_]
//
// Constraints:
// 0 <= nums.length <= 100
// 0 <= nums[i] <= 50
// 0 <= val <= 100

func removeElement(nums []int, val int) int {
	i := 0
	j := len(nums) - 1
	for i <= j {
		if nums[i] == val {
			nums[i], nums[j] = nums[j], nums[i]
			j--
			continue
		}
		i++
	}
	return i
}

func main() {
	fmt.Println(removeElement([]int{3, 2, 2, 3}, 3))
	fmt.Println(removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))
}
