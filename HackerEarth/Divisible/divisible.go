package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {

	var cnt int64
	fmt.Println("Enter no:")
	fmt.Scanf("%d", &cnt)

	arr := make([]int64, cnt)
	for i := int64(0); i < cnt; i++ {
		fmt.Scanf("%d", &arr[i])
	}

	index := len(arr) / 2

	var number string

	for i, val := range arr {
		val1 := i + 1
		if int64(val1) >= int64(index) {
			number += strconv.FormatInt(lastDigit(val)*int64(math.Pow10(len(arr)-1-i)), 10)
		} else {
			number += strconv.FormatInt(firstDigit(val)*int64(math.Pow10(len(arr)-1-i)), 10)
		}
	}
	var even, odd int
	for i := 0; i < len(number); i++ {
		if i%2 == 0 {
			x, _ := strconv.Atoi(string(number[i]))
			even += x
		} else {
			x, _ := strconv.Atoi(string(number[i]))
			odd += x
		}
	}

	if (even-odd)%11 == 0 {
		fmt.Println("OUI")
	} else {
		fmt.Println("NON")
	}

}

func firstDigit(n int64) int64 {
	for n >= 10 {
		n /= 10
	}
	return n
}

func lastDigit(n int64) int64 {
	for n >= 10 {
		n %= 10
	}
	return n
}
