package main

import "fmt"

func main() {
	var cycle int
	fmt.Print("Input Cycle:")
	fmt.Scanf("%d", &cycle)
	var output int

	for i := 0; i < cycle; i++ {
		var costOfgreenballon, costofpurpleballon int
		fmt.Print("Input Costs :")
		fmt.Scanf("%d", &costOfgreenballon)
		fmt.Scanf("%d", &costofpurpleballon)

		fmt.Print("the number of participants :")
		var participants int
		fmt.Scanf("%d", &participants)

		if (cycle < 1) || (cycle > 10) || (participants < 1) || (participants > 10) {
			fmt.Println("Validation failed")
			return
		}

		var g, p, g1, p1 int
		for j := 0; j < participants; j++ {
			var one, two int
			fmt.Scanf("%d", &one)
			fmt.Scanf("%d", &two)
			g = g + (one * costOfgreenballon)
			p = p + (two * costofpurpleballon)

			g1 = g1 + (one * costofpurpleballon)
			p1 = p1 + (two * costOfgreenballon)

			if g+p > g1+p1 {
				output = g1 + p1
			} else {
				output = g + p
			}
		}
		fmt.Println(output)
	}
}
