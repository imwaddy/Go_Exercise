package main

import "fmt"

func main() {
	var input int
	fmt.Println("Enter no:")
	fmt.Scanf("%d", &input)
	if input < 1 || input > 10 {
		return
	}
	fmt.Println(factorial(input))
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
