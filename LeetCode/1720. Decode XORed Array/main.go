package main

import (
	"fmt"
)

// 1720. Decode XORed Array
//
// Problem Statement:
// There is a hidden integer array arr of length n. It was encoded into
// array encoded of length n-1 where encoded[i] = arr[i] XOR arr[i+1].
// Given encoded and first element of arr (first), return original arr.
//
// Example 1:
// Input:  encoded = [1,2,3], first = 1
// Output: [1,0,2,1]
//
// Example 2:
// Input:  encoded = [6,2,7,3], first = 4
// Output: [4,2,0,7,4]
//
// Constraints:
// 2 <= n <= 10^4
// encoded.length == n - 1
// 0 <= encoded[i] <= 10^5
// 0 <= first <= 10^5

func decode(encoded []int, first int) []int {
	res := make([]int, len(encoded)+1)
	res[0] = first
	for i := 0; i < len(encoded); i++ {
		res[i+1] = res[i] ^ encoded[i]
	}
	return res
}

func main() {
	fmt.Println(decode([]int{1, 2, 3}, 1))
	fmt.Println(decode([]int{6, 2, 7, 3}, 4))
}
