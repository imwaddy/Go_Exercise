package main

import (
	"fmt"
)

// 448. Find All Numbers Disappeared in an Array
//
// Problem Statement:
// Given array nums of n integers where nums[i] in range [1, n], return
// array of all integers in range [1, n] that do not appear in nums.
//
// Example 1:
// Input:  nums = [4,3,2,7,8,2,3,1]
// Output: [5,6]
//
// Example 2:
// Input:  nums = [1,1]
// Output: [2]
//
// Constraints:
// n == nums.length
// 1 <= n <= 10^5
// 1 <= nums[i] <= n

// ---- not the optimal solution
// func findDisappearedNumbers(nums []int) []int {
// 	res := make([]int, 0)
// 	m := make(map[int]struct{}, 0)

// 	for _, num := range nums {
// 		m[num] = struct{}{}
// 	}

// 	for i := 1; i <= len(nums); i++ {
// 		_, exists := m[i]
// 		if !exists {
// 			res = append(res, i)
// 		}
// 	}

// 	return res
// }

// ---- optimal one

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func findDisappearedNumbers(nums []int) []int {
	for _, num := range nums {
		idx := abs(num) - 1
		if nums[idx] > 0 {
			nums[idx] = -nums[idx]
		}
	}

	res := make([]int, 0)
	for i, num := range nums {
		if num > 0 {
			res = append(res, i+1)
		}
	}
	return res
}

func main() {
	fmt.Println(findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1}))
	fmt.Println(findDisappearedNumbers([]int{1, 1}))
}
