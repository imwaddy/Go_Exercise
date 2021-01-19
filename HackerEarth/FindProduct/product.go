package main

import (
	"fmt"
	"math"
)

var (
	product = math.Pow10(9) + 7
)

func main() {

	var cnt int
	fmt.Println("size of the array:")
	fmt.Scanf("%d", &cnt)

	if cnt < 1 || cnt > int(math.Pow10(3)) {
		return
	}

	var arr = make([]int, cnt)

	var answer = 1
	for i := 0; i < cnt; i++ {

		fmt.Scanf("%d", &arr[i])

		if arr[i] < 1 || arr[i] > int(math.Pow10(3)) {
			return
		}

		answer = (answer * arr[i]) % int(product)
	}

	fmt.Println(answer)

}
