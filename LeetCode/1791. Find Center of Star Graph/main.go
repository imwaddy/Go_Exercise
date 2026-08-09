package main

import (
	"fmt"
)

// 1791. Find Center of Star Graph
//
// Problem Statement:
// There is an undirected star graph with n nodes labeled 1 to n. Star
// graph: one center node connected to all n-1 other nodes via n-1 edges.
// Given 2D array edges where edges[i] = [ui, vi] indicates edge between
// ui and vi, return center of star graph.
//
// Example 1:
// Input:  edges = [[1,2],[2,3],[4,2]]
// Output: 2
//
// Example 2:
// Input:  edges = [[1,2],[5,1],[1,3],[1,4]]
// Output: 1
//
// Constraints:
// 3 <= n <= 10^5
// edges.length == n - 1
// edges[i].length == 2
// 1 <= ui, vi <= n
// ui != vi
// Given edges represent valid star graph

func findCenter(edges [][]int) int {
	if edges[0][0] == edges[1][0] {
		return edges[0][0]
	} else if edges[0][0] == edges[1][1] {
		return edges[0][0]
	} else if edges[0][1] == edges[1][0] {
		return edges[0][1]
	} else {
		return edges[0][1]
	}

	return 0
}

func main() {
	fmt.Println(findCenter([][]int{{1, 2}, {2, 3}, {4, 2}}))
}
