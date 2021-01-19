package main

import (
	"fmt"
	"strings"
)

func main() {

	var noOfGrids int
	fmt.Print("Input no of splits:")
	fmt.Scanf("%d", &noOfGrids)
	if (noOfGrids < 1) || (noOfGrids > 20) {
		fmt.Println("NO")
		return
	}

	var village string
	fmt.Print("Input village:")
	fmt.Scanf("%s", &village)

	var isAllowed = strings.Contains(village, "HH")
	if isAllowed {
		fmt.Println("NO")
		return
	}

	village = strings.ReplaceAll(village, ".", "B")

	fmt.Println("YES")
	fmt.Println(village)

}
