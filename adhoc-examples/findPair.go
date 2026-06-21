package main

import "fmt"

func findPairs() {
	nums := []int{1, 2, 3, 4, 5, 6}
	target := 7
	result := findPairsinternal(nums, target)
	fmt.Printf("Pairs with sum equal to %d are: %v\n", target, result)
}

func findPairsinternal(nums []int, target int) [][]int {
	numSet := make(map[int]bool)
	pairs := [][]int{}

	for _, num := range nums {
		complement := target - num
		if numSet[complement] {
			pairs = append(pairs, []int{complement, num})
		}
		numSet[num] = true
	}

	return pairs
}

// Sofa
// Cot
// Cupboard,
// Washing m/c
// Fridge
// Dressing Table
// Rack + Dishes
// Mixer
// Cooker
