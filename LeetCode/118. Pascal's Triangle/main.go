package main

import "fmt"

// 118. Pascal's Triangle
//
// Problem Statement:
// Given integer numRows, return first numRows of Pascal's triangle.
// Each number is sum of two numbers directly above it.
//
// Example 1:
// Input:  numRows = 5
// Output: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]
//
// Example 2:
// Input:  numRows = 1
// Output: [[1]]
//
// Constraints:
// 1 <= numRows <= 30

// copied answer
func generate(numRows int) [][]int {
	result := make([][]int, numRows)

	for i := 0; i < numRows; i++ {
		result[i] = make([]int, i+1)

		// First and last values in each row are 1
		result[i][0] = 1
		result[i][i] = 1

		// Fill inner values using the previous row
		for j := 1; j < i; j++ {
			result[i][j] = result[i-1][j-1] + result[i-1][j]
		}
	}

	return result
}

func main() {
	fmt.Println(generate(5))
	fmt.Println(generate(1))
}
