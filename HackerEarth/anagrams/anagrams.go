package main

import (
	"fmt"
)

func main() {
	fmt.Println(anagramCount("cde", "abc"))
}

func anagramCount(s1, s2 string) int {
	var cnt int
	if len(s1) >= len(s2) {
		mapstr := mapFirstString(s1)
		for _, i := range s2 {
			_, ok := mapstr[string(i)]
			if ok {
				cnt += 2
			}
		}
	} else {
		mapstr := mapFirstString(s2)
		for _, i := range s1 {
			_, ok := mapstr[string(i)]
			if !ok {
				cnt++
			}
		}
	}
	return (len(s1) + len(s2)) - cnt

}

func mapFirstString(s string) map[string]string {
	first := make(map[string]string, 0)
	for _, i := range s {
		first[string(i)] = string(i)
	}
	return first
}
