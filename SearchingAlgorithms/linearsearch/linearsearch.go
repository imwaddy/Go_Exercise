package main

import "fmt"

// TODO: Convert it into generics

func main() {
	// Sample array for testing linearSearch
	arr := []int{22, 33, 55, 66, 66, 7711, 243, 12132, 4355, 35346}

	// Calling linearSearch to find the number 55 in the array
	found := linearSearch(arr, 55)

	// Printing the result
	fmt.Println("Number Found:", found)
}

// linearSearch function performs a linear search on an integer array
func linearSearch(arr []int, key int) bool {
	// Iterating over each element in the array
	for _, elem := range arr {
		// Checking if the current element matches the target key
		if elem == key {
			// Return true if the key is found
			return true
		}
	}

	// Return false if the key is not found in the array
	return false
}
