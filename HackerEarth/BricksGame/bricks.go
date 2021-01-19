package main

import "fmt"

func main() {
	var no int
	fmt.Print("Enter no:")
	fmt.Scanf("%d", &no)

	if no < 1 && no > 10000 {
		return
	}
	cnt, previ := 0, 0
	for i := 0; ; i++ {
		previ = i
		if i%2 == 0 {
			cnt++
			if cnt == no {
				fmt.Println("Patlu")
				return
			}
		} else {
			cnt += (previ * 2)
			if cnt == no || cnt > no {
				fmt.Println("Motu")
				return
			}
		}
		fmt.Printf("Count=%d and i=%d", cnt, i)
	}
}
