package main

import (
	"fmt"
	"math"
)

func main() {
	var cnt int64
	fmt.Scanf("%d", &cnt)
	if cnt < int64(1) || cnt > int64(100) {
		return
	}

	for i := int64(0); i < cnt; i++ {

		var no int64
		fmt.Scanf("%d", &no)
		if cnt < int64(2) || cnt > int64(math.Pow10(1000000)) {
			return
		}
		fmt.Println("In loop 1")

		for no != int64(1) {
			fmt.Println("In loop 22")
			if no%2 == 0 {
				no = no / 2
			} else if no%2 == 1 {
				no = (no * 3) + 1
			}
		}
		fmt.Println("In loop 33")
		if no == 1 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}

	}
}
