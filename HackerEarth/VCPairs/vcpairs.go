package main

import "fmt"

func main() {

	var testCases int
	fmt.Println("Enter test cases:")
	fmt.Scanf("%d", &testCases)

	for i := 0; i < testCases; i++ {
		var cnt int
		fmt.Println("Enter no of chars:")
		fmt.Scanf("%d", &cnt)

		var str string
		fmt.Println("Enter string:")
		fmt.Scanf("%s", &str)

		var count int

		for k := 0; k < cnt-1; k++ {
			if !isVowel(string(str[k])) && isVowel(string(str[k+1])) {
				count++
			}
		}

		fmt.Println(count)

	}

}

func isVowel(data string) bool {
	if data == "a" || data == "e" || data == "i" || data == "o" || data == "u" {
		return true
	}
	return false
}
