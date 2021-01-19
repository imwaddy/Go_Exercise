package main

import "fmt"

func main() {
	a := 0

	sticks := map[int]int{
		0: 6, 1: 2, 2: 5, 3: 5, 4: 4, 5: 5, 6: 6, 7: 3, 8: 7, 9: 6,
	}
	cnt := 0
	for a != 0 {
		rem := a % 10
		val, ok := sticks[rem]
		if ok {
			cnt += val
		}
		a = a / 10
	}
	no := "1"
	for _, val := range sticks {
		a *= 10
		no += "1"
		cnt -= val
	}
	fmt.Println(no)
}
