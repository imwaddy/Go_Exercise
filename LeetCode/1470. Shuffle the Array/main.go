package main

import (
	"fmt"
)

// 1470. Shuffle the Array
//
// Problem Statement:
// Given array nums of 2n elements in form [x1,x2,...,xn,y1,y2,...,yn],
// return array in form [x1,y1,x2,y2,...,xn,yn].
//
// Example 1:
// Input:  nums = [2,5,1,3,4,7], n = 3
// Output: [2,3,5,4,1,7]
//
// Example 2:
// Input:  nums = [1,2,3,4,4,3,2,1], n = 4
// Output: [1,4,2,3,3,2,4,1]
//
// Constraints:
// 1 <= n <= 500
// nums.length == 2n
// 1 <= nums[i] <= 10^3

func shuffle(nums []int, n int) []int {
	res := make([]int, len(nums))

	for i := 0; i < n; i++ {
		res[2*i] = nums[i]
		res[2*i+1] = nums[n+i]
	}

	return res
}

// shuffleInPlace: O(1) extra space, encode two values in one int.
// nums[i] <= 1000, so base 1001 fits both x and y without collision.
func shuffleInPlace(nums []int, n int) []int {
	const base = 1001

	// encode: nums[i] = x*base + y, pairing x_i with y_i (index n+i)
	for i := 0; i < n; i++ {
		nums[i] = nums[i]*base + nums[n+i]
	}

	// decode from back to front, writing into final positions
	last := len(nums) - 1
	for i := n - 1; i >= 0; i-- {
		y := nums[i] % base
		x := nums[i] / base
		nums[last] = y
		nums[last-1] = x
		last -= 2
	}

	return nums
}

func main() {
	fmt.Println(shuffle([]int{2, 5, 1, 3, 4, 7}, 3))
	fmt.Println(shuffleInPlace([]int{2, 5, 1, 3, 4, 7}, 3))
}
