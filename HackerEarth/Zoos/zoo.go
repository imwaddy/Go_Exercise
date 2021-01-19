package main

import (
	"fmt"
	"strings"
)

func main() {

	var word string
	fmt.Println("Enter string:")
	fmt.Scanf("%s", &word)

	if len(word) < 1 || len(word) > 20 {
		return
	}

	x := strings.Count(word, "z")
	if strings.Count(word, "o") == x*2 {
		fmt.Println("Yes")
		return
	}
	fmt.Println("No")
}
