// Author: Mayur Wadekar

package main

import "fmt"

/*
	Longest Consecutive Sequence

	// input1
	input1: Length of an array
	input2: {100, 4, 5, 80, 30, 24}
	output: Longest Consecutive Sequence {4,5} so length is 2

	// input2
	input1: Length of an array
	input2: {100, 4, 200, 80, 30, 24}
	output: Longest Consecutive Sequence {} so length is 1

*/

var lengthOfArr int

// main function
func main() {

	// length of an array
	fmt.Print("Enter array length=")
	fmt.Scanf("%d", &lengthOfArr)

	// array
	arr := make([]int, 5)

	// Take array input from user
	for i := 0; i < 5; i++ {
		fmt.Print("Enter ", i, " array elements=")
		fmt.Scanf("%d", &arr[i])
	}

	// reassign sorted array
	arr = bubbleSort(arr)

	fmt.Println("Longest consecutive sequence length is ", CheckSequence(arr))
}

// CheckSequence checks for core logic to check for Consecutive Sequence
func CheckSequence(arr []int) int {

	length, Val := 0, 1
	for i := 1; i < len(arr); i++ {
		if arr[i]-1 == arr[i-1] {
			Val++
		} else if arr[i] != arr[i-1] {
			Val = 1
		}

		if length < Val {
			length = Val
		}
	}

	// returning parameter
	return length

}

// bubbleSort sorts element using bubble sort algorithm
func bubbleSort(a []int) []int {

	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}

	return a

}
