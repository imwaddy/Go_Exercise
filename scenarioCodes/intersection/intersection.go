// Given two integer arrays nums1 and nums2, return an array of their intersection.
// Each eleme the result must appear as many times as it shows in both arrays,
// and you may return the resul any order.

// Example 1:

// Input: nums1 = [1,2,2,1], nums2 = [2,2] Output: [2,2] Example 2:

// Input: nums1=[4,9,5], nums2= [9,4,9,8,4] Output: [4,9] Explanation: [9,4] is also accepted.

package main

import "fmt"

func main() {

	result := getIntersection([]int{9, 4, 9, 8, 4}, []int{4, 9, 5})

	fmt.Println("Result ", result)
}

func getIntersection(arr1, arr2 []int) []int {
	m := make(map[int]int)

	for _, r := range arr1 {
		m[r]++
	}

	var intersection []int

	for _, r := range arr2 {
		if m[r] > 0 {
			intersection = append(intersection, r)
			m[r]--
		}
	}
	return intersection

}
