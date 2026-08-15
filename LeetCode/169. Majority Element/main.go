package main

import "fmt"

func check(name string, got int, want int) {
	if got == want {
		fmt.Printf("%s PASS got=%v\n", name, got)
	} else {
		fmt.Printf("%s FAIL got=%v want=%v\n", name, got, want)
	}
}

// 169. Majority Element
//
// Problem Statement:
// Given an array nums of size n, return the majority element. The majority
// element is the element that appears more than n/2 times. You may assume
// the majority element always exists in the array.
//
// Example 1:
// Input:  nums = [3,2,3]
// Output: 3
//
// Example 2:
// Input:  nums = [2,2,1,1,1,2,2]
// Output: 2
//
// Constraints:
// n == nums.length
// 1 <= n <= 5 * 10^4
// -10^9 <= nums[i] <= 10^9

func majorityElement(nums []int) int {
	a := []int{nums[0], 1}
	for i := 1; i <= len(nums)-1; i++ {
		if nums[i] != a[0] {
			a[1] -= 1
		} else {
			a[1] += 1
		}

		if a[1] < 0 {
			a[0] = nums[i]
			a[1] = 1
		}
	}
	return a[0]
}

func main() {
	check("[1]", majorityElement([]int{3, 2, 3}), 3)
	check("[2]", majorityElement([]int{2, 2, 1, 1, 1, 2, 2}), 2)
	check("[3]", majorityElement([]int{1}), 1)
	check("[4]", majorityElement([]int{-1, -1, 2}), -1)
	check("[5]", majorityElement([]int{1, 2, 1, 1, 2}), 1)
}
