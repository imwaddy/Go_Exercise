package main

import "fmt"

func main() {
	Palidrome("fnjzxnxnjplfwzowfdrhrvhegkmoncbkembjoudteqchjwqfzlofyflkmxnooasxulwofjzknthqqxgshvwxdvhdnlzjzdjdiifg")
}

func Palidrome(s string) {
	if len(s) == 1 {
		fmt.Println("YES")
	}
	if len(s) == 2 && s[0] == s[1] {
		fmt.Println("YES")
	} else if len(s) == 2 && s[0] != s[1] {
		fmt.Println("NO")
	}
	for i, j := 0, len(s)-1; i <= j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			fmt.Println("NO")
			return
		}
	}
	fmt.Println("YES")
}
