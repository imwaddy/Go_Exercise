// Author: Mayur Wadekar

package main

import (
	"fmt"
	"strings"
)

func main() {
	// variable declaration
	var stringToCheck string

	// Get user input
	fmt.Print("Please Enter string==")
	fmt.Scanf("%s", &stringToCheck)

	// convert to lower case
	stringToCheck = strings.TrimSpace(strings.ToLower(stringToCheck))

	// default consideration that string is palindrome
	isPalindrome := true

	// loop through variable
	i := 0
	j := len(stringToCheck) - 1

	for {

		// If string having only single character
		if i == j || len(stringToCheck) == 1 {
			break
		}

		// If string having only two character
		if len(stringToCheck) == 2 && stringToCheck[i] == stringToCheck[j] {
			break
		}

		if string(stringToCheck[i]) != string(stringToCheck[j]) {
			isPalindrome = false
			break
		}
		i++
		j--
	}

	// Conclusion
	if isPalindrome {
		fmt.Println("String Is Palindrome")
	} else {
		fmt.Println("String Is Not Palindrome")
	}

}
