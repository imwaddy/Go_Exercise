package main

import (
	"fmt"
	"math"
)

var (
	floors = 7
)

func main() {

	var seq int
	fmt.Print("Enter loop:")
	fmt.Scanf("%d", &seq)

	var liftA, liftB = 0, 7

	for i := 0; i < seq; i++ {

		var calledFromFloor int
		fmt.Print("Enter floor:")
		fmt.Scanf("%d", &calledFromFloor)

		if math.Abs(float64(liftA-calledFromFloor)) <= math.Abs(float64(liftB-calledFromFloor)) {
			fmt.Println("A")
			liftA = calledFromFloor
		} else {
			fmt.Println("B")
			liftB = calledFromFloor
		}
	}
}
