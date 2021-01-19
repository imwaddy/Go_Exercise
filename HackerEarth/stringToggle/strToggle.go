package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	fmt.Print("Enter string:")
	fmt.Scanf("%s", &str)
	if len(str) < 1 || len(str) > 100 {
		return
	}

	var data string

	for i := 0; i < len(str); i++ {
		fmt.Println(str[i])
		if str[i] >= 97 && str[i] <= 122 {
			data += strings.ToUpper(string(str[i]))
		} else if str[i] >= 65 && str[i] <= 90 {
			data += strings.ToLower(string(str[i]))
		}
	}

	fmt.Println(data)
}
