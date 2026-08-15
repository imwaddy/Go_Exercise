package main

import "fmt"

func check(name string, got int, want int) {
	if got == want {
		fmt.Printf("%s PASS got=%v\n", name, got)
	} else {
		fmt.Printf("%s FAIL got=%v want=%v\n", name, got, want)
	}
}

// 455. Assign Cookies
//
// Problem Statement:
// Each child i has greed factor g[i] (min cookie size to be content). Each
// cookie j has size s[j]. A child is content if given cookie j with
// s[j] >= g[i]. Each child gets at most one cookie. Maximize number of
// content children.
//
// Example 1:
// Input:  g = [1,2,3], s = [1,1]
// Output: 1
//
// Example 2:
// Input:  g = [1,2], s = [1,2,3]
// Output: 2
//
// Constraints:
// 1 <= g.length <= 3 * 10^4
// 0 <= s.length <= 3 * 10^4
// 1 <= g[i], s[j] <= 2^31 - 1

func findContentChildren(g []int, s []int) int {
	// TODO: implement
	return 0
}

func main() {
	check("[1]", findContentChildren([]int{1, 2, 3}, []int{1, 1}), 1)
	check("[2]", findContentChildren([]int{1, 2}, []int{1, 2, 3}), 2)
	check("[3]", findContentChildren([]int{1, 2, 3}, []int{}), 0)
	check("[4]", findContentChildren([]int{}, []int{1, 2}), 0)
	check("[5]", findContentChildren([]int{10, 9, 8, 7}, []int{5, 6, 7, 8}), 2)
}
