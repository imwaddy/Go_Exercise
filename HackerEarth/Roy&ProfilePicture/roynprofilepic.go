package main

import "fmt"

func main() {
	var length int
	fmt.Println("Enter Length ")
	fmt.Scanf("%d", &length)
	if length < 1 || length > 10000 {
		return
	}

	var noofphotos int
	fmt.Println("Enter no of photos ")
	fmt.Scanf("%d", &noofphotos)
	if noofphotos < 1 || noofphotos > 1000 {
		return
	}

	for i := 0; i < noofphotos; i++ {
		var H, W int

		fmt.Println("Enter Height (Space) Width ")
		fmt.Scanf("%d%d", &H, &W)

		if H < length || W < length {
			fmt.Println("UPLOAD ANOTHER")
			continue
		} else if H > length || W > length {
			fmt.Println("CROP IT")
			continue
		} else if H == length || W == length {
			fmt.Println("ACCEPTED")
			continue
		}
	}
}
