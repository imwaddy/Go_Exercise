package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Println(getLucky("fleyctuuajsr", 5))
}

func getLucky(s string, k int) int {
	alphabetMap := make(map[rune]string)

	for i := 'a'; i <= 'z'; i++ {
		alphabetMap[i] = strconv.Itoa(int(i - 'a' + 1))
	}

	var result int
	s = convertToStringTypeNumber(s, alphabetMap)

	for i := 0; i < k; i++ {
		no, _ := strconv.Atoi(s)
		result = sumOfDigits(no)
		s = convertToStringTypeNumber(strconv.Itoa(result), alphabetMap)

		fmt.Println(s)
	}

	return result
}

func sumOfDigits(n int) int {
	var sum int
	for n != 0 {
		rem := n % 10
		sum = sum + rem
		n = n / 10
	}
	return sum
}

func convertToStringTypeNumber(s string, alphabetMap map[rune]string) string {
	var num string
	for _, s1 := range s {
		num += string(alphabetMap[s1])
	}
	return num
}
