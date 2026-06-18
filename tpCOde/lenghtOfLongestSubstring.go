package main

import "fmt"

func lenghtOfLongestSubstring() {

	str1 := "ABDEFGABEF"

	lastOccurrence := make(map[byte]int)
	var start, maxLength int

	for i, _ := range str1 {

		idx, found := lastOccurrence[str1[i]]
		if found && idx >= start {
			start = idx + 1
		}
		lastOccurrence[str1[i]] = i

		if i-start+1 > maxLength {
			fmt.Println("ok ok ok ok ", start+1, i)
			maxLength = i - start + 1
		}
	}

	// fmt.Println("Maxlenght ", str1[start+1:len(str1)-1-maxLength])

}
