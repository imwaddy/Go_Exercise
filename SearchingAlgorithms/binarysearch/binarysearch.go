package main

import (
	"fmt"
	"sort"
)

// TODO: Convert it into generics

func main() {
	// Sample array for testing BinarySearch
	arr := []int{1, 2, 3, 4, 5}

	// Sorting the array as BinarySearch typically requires a sorted array
	sort.Ints(arr)

	// Calling BinarySearch to find the number 1 in the array
	found := BinarySearch(arr, 6)

	// Printing the result
	fmt.Println("Number Found:", found)
}

// BinarySearch performs a binary search on a sorted integer array
func BinarySearch(arr []int, e int) bool {

	// If array of size zero then return false
	if len(arr) == 0 {
		return false
	}

	// Calculating the middle index
	mid := len(arr) / 2

	// Checking if the middle element is equal to the target
	if arr[mid] == e {
		return true
	} else if e < arr[mid] {
		// If the target is less than the middle element, recursively search the left half
		return BinarySearch(arr[0:mid], e)
	} else if e > arr[mid] {
		// If the target is greater than the middle element, recursively search the right half
		return BinarySearch(arr[mid+1:], e)
	}

	// If the target is not found, return false
	return false
}
