package main

import (
	"fmt"
	"sync"
)

// 1672. Richest Customer Wealth
//
// Problem Statement:
// Given m x n grid accounts where accounts[i][j] is amount of money
// customer i has in bank j. Return wealth of richest customer (max sum
// of a single row).
//
// Example 1:
// Input:  accounts = [[1,2,3],[3,2,1]]
// Output: 6
//
// Example 2:
// Input:  accounts = [[1,5],[7,3],[3,5]]
// Output: 10
//
// Example 3:
// Input:  accounts = [[2,8,7],[7,1,3],[1,9,5]]
// Output: 17
//
// Constraints:
// m == accounts.length
// n == accounts[i].length
// 1 <= m, n <= 50
// 1 <= accounts[i][j] <= 100

func total(nums []int, wg *sync.WaitGroup, res chan int) {
	defer wg.Done()
	var sum int
	for _, no := range nums {
		sum += no
	}
	res <- sum
}

func maximumWealth(accounts [][]int) int {
	var max = 0
	wg := sync.WaitGroup{}
	res := make(chan int)

	for _, acc := range accounts {
		wg.Add(1)
		go total(acc, &wg, res)
	}

	go func() {
		wg.Wait()
		close(res)
	}()

	for r := range res {
		if r > max {
			max = r
		}
	}

	return max
}

func main() {
	fmt.Println(maximumWealth([][]int{{1, 2, 3}, {3, 2, 1}}))
}
