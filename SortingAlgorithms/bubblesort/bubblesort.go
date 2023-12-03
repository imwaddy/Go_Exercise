package main

import "fmt"

func main() {
	// Example usage of bubbleSort function
	inputArray := []int{9, 5, 4, 3, 21, 3}
	bubbleSort(inputArray)
	fmt.Println(inputArray)
}

// bubbleSort sorts an integer slice in ascending order using the bubble sort algorithm.
// It takes an integer slice as input and sorts it in place.
func bubbleSort(arr []int) {
	n := len(arr)

	// Traverse through all elements in the array
	for i := 0; i < n-1; i++ {
		// Last i elements are already sorted, so we don't need to check them
		for j := 0; j < n-i-1; j++ {
			// Swap if the element found is greater than the next element
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}
