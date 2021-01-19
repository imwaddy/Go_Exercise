package main

import "fmt"

func main() {

	var str string
	fmt.Println("Enter input:")
	fmt.Scanf("%s", &str)

	var x, y int
	for i := 0; i < len(str); i++ {
		if string(str[i]) == "L" {
			x--
		} else if string(str[i]) == "R" {
			x++
		} else if string(str[i]) == "U" {
			y++
		} else {
			y--
		}
	}
	fmt.Println(x, y)
}
