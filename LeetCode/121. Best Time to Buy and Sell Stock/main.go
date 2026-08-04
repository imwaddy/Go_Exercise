package main

import "fmt"

// 121. Best Time to Buy and Sell Stock
//
// Problem Statement:
// Given array prices where prices[i] is price of stock on day i, you want
// to maximize profit by choosing single day to buy and different later day
// to sell. Return max profit achievable. If no profit possible, return 0.
//
// Example 1:
// Input:  prices = [7,1,5,3,6,4]
// Output: 5
// (buy day 2 price=1, sell day 5 price=6, profit=5)
//
// Example 2:
// Input:  prices = [7,6,4,3,1]
// Output: 0
// (prices only fall, no profit possible)
//
// Constraints:
// 1 <= prices.length <= 10^5
// 0 <= prices[i] <= 10^4
//
// Hint:
// One pass. Track minimum price seen so far as you scan left to right.
// At each day, check profit if sold today (price - minSoFar), keep max.

func MaxProfit(prices []int) int {
	maxProfit := 0
	min := prices[0]

	for i := 1; i <= len(prices)-1; i++ {
		if min > prices[i] {
			min = prices[i]
		}
		if maxProfit < (prices[i] - min) {
			maxProfit = prices[i] - min
		}

	}

	return maxProfit
}

func main() {
	tests := []struct {
		input []int
		want  int
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 5},
		{[]int{7, 6, 4, 3, 1}, 0},
		{[]int{1}, 0},
		{[]int{2, 4, 1}, 2},
		{[]int{3, 3, 3, 3}, 0},
		{[]int{1, 2}, 1},
	}

	for _, tc := range tests {
		got := MaxProfit(tc.input)
		status := "PASS"
		if got != tc.want {
			status = "FAIL"
		}
		fmt.Printf("[%s] got=%v want=%v prices=%v\n", status, got, tc.want, tc.input)
	}
}
