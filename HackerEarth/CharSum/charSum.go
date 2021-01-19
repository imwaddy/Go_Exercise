package main

import "fmt"

func main() {
	values := makeStringMap()

	var input string
	fmt.Print("Enter string:")
	fmt.Scanf("%s", &input)
	if len(input) < 1 || len(input) > 100 {
		return
	}

	var cnt int
	for _, val := range input {
		mapVal, _ := values[string(val)]
		cnt += mapVal
	}
	fmt.Println(cnt)
}

func makeStringMap() map[string]int {
	val := make(map[string]int, 0)
	a := 97
	for i := 1; i <= 26; i++ {
		val[string(a)] = i
		a++
	}
	return val
}
