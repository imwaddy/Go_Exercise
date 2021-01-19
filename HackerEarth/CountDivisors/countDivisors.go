package main

import "fmt"

func main() {
	var first, second, third int
	fmt.Println("Enter 3 numbers:")
	fmt.Scanf("%d%d%d", &first, &second, &third)

	if first < 1 || second < 1 || second > 1000 || first > 1000 || third < 1 || third > 1000 {
		return
	}

	var count int

	for i := first; i <= second; i++ {
		if i%third == 0 {
			count++
		}
	}
	fmt.Println(count)
}
