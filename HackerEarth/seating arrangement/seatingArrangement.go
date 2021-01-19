package main

import "fmt"

const (
	maxSeats    = 12
	decrementBy = 2
)

func main() {
	var cycle int
	fmt.Print("Input Cycle:")
	fmt.Scanf("%d", &cycle)
	if (cycle < 1) || (cycle > 100000) {
		return
	}

	seatingList := []string{
		1: "WS", 2: "MS", 3: "AS",
		4: "AS", 5: "MS", 6: "WS",
		7: "WS", 8: "MS", 9: "AS",
		10: "AS", 11: "MS", 0: "WS",
	}

	for i := 0; i < cycle; i++ {
		var input int
		fmt.Println("Enter seat number: ")
		fmt.Scanf("%d", &input)
		if (input < 1) || (input > 108) {
			return
		}
		var index = getMaxSeat(input)
		input += calculateSeat(index)
		fmt.Println(input, " ", seatingList[index])
	}
}

func getMaxSeat(input int) int {
	return input % maxSeats
}

func calculateSeat(index int) int {

	tempMaxSeat := maxSeats - 1

	if index == 0 {
		return -tempMaxSeat
	}

	for i := 0; i < index-1; i++ {

		tempMaxSeat -= decrementBy

	}
	return tempMaxSeat
}
