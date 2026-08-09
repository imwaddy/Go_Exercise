package main

import (
	"fmt"
)

// 1108. Defanging an IP Address
//
// Problem Statement:
// Given valid IPv4 address, return defanged version (replace every "."
// with "[.]").
//
// Example 1:
// Input:  address = "1.1.1.1"
// Output: "1[.]1[.]1[.]1"
//
// Example 2:
// Input:  address = "255.100.50.0"
// Output: "255[.]100[.]50[.]0"
//
// Constraints:
// address is valid IPv4 address

func defangIPaddr(address string) string {
	// return strings.Replace(address, ".", "[.]", -1)
	ipAddr := ""
	for _, ch := range address {
		if ch == '.' {
			ipAddr += "[.]"
		} else {
			ipAddr += string(ch)
		}
	}
	return ipAddr
}

func main() {
	fmt.Println(defangIPaddr("1.1.1.1"))
}
