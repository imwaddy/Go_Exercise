// Author; Mayur Wadekar
/*
	This is simple goroutine example
*/

package main

import "fmt"

// printNumber: simply prints numbers
func printNumber(n int) {
	for i := 0; i < n; i++ {
		fmt.Println(n, ":", i)
	}
}

// main function
func main() {
	go printNumber(10)
	// Wait to finish goroutine
	var input string
	fmt.Scanln(&input)
}
