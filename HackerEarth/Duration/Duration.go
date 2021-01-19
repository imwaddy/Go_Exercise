package main

import "fmt"

func main() {

	var count int
	fmt.Println("Enter loop:")
	fmt.Scanf("%d", &count)

	for i := 0; i < count; i++ {
		var SH, SM, EH, EM int
		fmt.Println("Enter SH, SM, EH, EM: ")
		fmt.Scanf("%d%d%d%d", &SH, &SM, &EH, &EM)

		var stime = (60 * SH) + SM

		var etime = (60 * EH) + EM

		var finalMinutes = (etime - stime) % 60

		var finalHours = (etime - stime) / 60

		fmt.Println(finalHours, finalMinutes)
	}
}
