// Author: Mayur Wadekar
package main

import "fmt"

// main function
func main() {
	fmt.Println("Called stackAnalysis", stackAnalysis())
}

// stackAnalysis returns int variable
//go:noinline
func stackAnalysis() int {
	data := 55
	return data
}
