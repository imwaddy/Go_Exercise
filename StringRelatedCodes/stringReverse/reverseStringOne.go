// Author: Mayur Wadekar

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
input:  "Hi my name is mayur"
output: "Ih Ym Eman Si Ruyam"
*/

// main function
func main() {

	// variable declaration
	var str string

	// Get user input
	fmt.Print("Please Enter string==")
	// fmt.Scanf("%s", &str)
	reader := bufio.NewReader(os.Stdin)

	str, err := reader.ReadString('\n')
	if err != nil {
		fmt.Errorf("Error while reading string")
		os.Exit(0)
	}

	if str == "" {
		fmt.Errorf("Blank string is not allowed")
		os.Exit(0)
	}

	// printing final output with title case
	fmt.Println(GetReversedString(str))

}

// GetReversedString returns whole concatenated string which is reverse
func GetReversedString(str string) string {

	// triming spaces from left and right from string and splitting string through space
	arrStr := strings.Split(strings.TrimSpace(str), " ")

	// taking blank string for generating to new string
	str = ""

	// array through each word and reverse every word
	for _, s := range arrStr {
		s := Reverse(s)
		str += s + " "
	}

	// printing final output with title case
	return strings.TrimSpace(strings.Title(strings.ToLower(str)))
}

// Reverse function return reverse string
func Reverse(s string) string {
	n := len(s)
	runes := make([]rune, n)
	for _, rune := range s {
		n--
		runes[n] = rune
	}

	return string(runes[n:])
}
