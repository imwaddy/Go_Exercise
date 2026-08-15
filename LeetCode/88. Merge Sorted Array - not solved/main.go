package main

import (
	"fmt"
	"reflect"
)

func check(name string, got []int, want []int) {
	if reflect.DeepEqual(got, want) {
		fmt.Printf("%s PASS got=%v\n", name, got)
	} else {
		fmt.Printf("%s FAIL got=%v want=%v\n", name, got, want)
	}
}

// 88. Merge Sorted Array
//
// Problem Statement:
// You are given two sorted integer arrays nums1 and nums2, and integers m
// and n representing number of elements in nums1 and nums2 respectively.
// nums1 has length m+n, with last n elements set to 0 (placeholder space).
// Merge nums2 into nums1 as one sorted array, in-place.
//
// Example 1:
// Input:  nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
// Output: [1,2,2,3,5,6]
//
// Example 2:
// Input:  nums1 = [1], m = 1, nums2 = [], n = 0
// Output: [1]
//
// Constraints:
// nums1.length == m + n
// nums2.length == n
// 0 <= m, n <= 200
// 1 <= m + n <= 200

func merge(nums1 []int, m int, nums2 []int, n int) {
	return
}

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	merge(nums1, 3, []int{2, 5, 6}, 3)
	check("[1]", nums1, []int{1, 2, 2, 3, 5, 6})

	nums2 := []int{1}
	merge(nums2, 1, []int{}, 0)
	check("[2]", nums2, []int{1})

	nums3 := []int{0}
	merge(nums3, 0, []int{1}, 1)
	check("[3]", nums3, []int{1})

	nums4 := []int{1, 2, 0, 0}
	merge(nums4, 2, []int{5, 6}, 2)
	check("[4]", nums4, []int{1, 2, 5, 6})

	nums5 := []int{4, 5, 6, 0, 0, 0}
	merge(nums5, 3, []int{1, 2, 3}, 3)
	check("[5]", nums5, []int{1, 2, 3, 4, 5, 6})

	nums6 := []int{2, 4, 6, 0, 0, 0}
	merge(nums6, 3, []int{2, 4, 6}, 3)
	check("[6]", nums6, []int{2, 2, 4, 4, 6, 6})
}
