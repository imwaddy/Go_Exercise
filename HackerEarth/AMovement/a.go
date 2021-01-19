package main

import (
	"fmt"
	"math"
)

func main() {

	var cnt int
	fmt.Println("Enter no:")
	fmt.Scanf("%d", &cnt)
	if cnt < 1 || float64(cnt) > math.Pow10(6) {
		return
	}

	var steps, carry = 5, 0

	for cnt != 0 {
		if cnt < 0 || steps < 0 {
			break
		}
		if steps < cnt {
			cnt -= steps
			carry++
		} else {
			cnt -= steps
			carry++
		}
	}
	fmt.Println(carry)
}
