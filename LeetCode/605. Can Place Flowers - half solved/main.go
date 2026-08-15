package main

import "fmt"

func check(name string, got bool, want bool) {
	if got == want {
		fmt.Printf("%s PASS got=%v\n", name, got)
	} else {
		fmt.Printf("%s FAIL got=%v want=%v\n", name, got, want)
	}
}

// 605. Can Place Flowers
//
// Problem Statement:
// You have a long flowerbed, some plots planted (1) and some empty (0). No
// two flowers can be planted in adjacent plots. Given flowerbed and integer
// n, return true if n new flowers can be planted without violating the rule.
//
// Example 1:
// Input:  flowerbed = [1,0,0,0,1], n = 1
// Output: true
//
// Example 2:
// Input:  flowerbed = [1,0,0,0,1], n = 2
// Output: false
//
// Constraints:
// 1 <= flowerbed.length <= 2 * 10^4
// flowerbed[i] is 0 or 1
// there are no two adjacent flowers in flowerbed
// 0 <= n <= flowerbed.length

func canPlaceFlowers(flowerbed []int, n int) bool {
	for i := 1; i <= len(flowerbed)-2; i++ {
		if flowerbed[i] == 0 && flowerbed[i-1] == 0 && flowerbed[i+1] == 0 {
			n--
			flowerbed[i] = 1
		}
	}
	return n == 0
}

func main() {
	check("[1]", canPlaceFlowers([]int{1, 0, 0, 0, 1}, 1), true)
	check("[2]", canPlaceFlowers([]int{1, 0, 0, 0, 1}, 2), false)
	check("[3]", canPlaceFlowers([]int{0}, 1), true)
	check("[4]", canPlaceFlowers([]int{0, 0, 0}, 2), true)
	check("[5]", canPlaceFlowers([]int{1}, 0), true)
	check("[6]", canPlaceFlowers([]int{1}, 1), false)
}
