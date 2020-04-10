// Author: Mayur Wadekar
package main

import "fmt"

// main function
func main() {
	fmt.Println("Called heapAnalysis", heapAnalysis())
}

// heapAnalysis returns *int pointer
//go:noinline
func heapAnalysis() *int {
	data := 55
	return &data
}
