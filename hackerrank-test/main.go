package main

import (
	"fmt"
	"sort"
)

func main() {

	a := []string{"a", "abc", "ab", "b"}

	fmt.Println(check(a))

}

func check(arr []string) []string {
	var result []string
	m := make(map[int][]string)
	mapp := []int{}
	for _, r1 := range arr {
		if !checkIfExists(mapp, len(r1)) {
			mapp = append(mapp, len(r1))
		}

		r, ok := m[len(r1)]
		if ok {
			r = append(r, r1)
			m[len(r1)] = r
		} else {
			rr := []string{r1}
			m[len(r1)] = rr
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(mapp)))

	fmt.Println("MAPP ", mapp)

	for _, val := range mapp {
		v, _ := m[val]
		if len(v) > 1 {
			sort.Strings(v)
		}
		result = append(result, v...)
	}

	return result
}

func checkIfExists(mapp []int, len int) bool {

	for _, a := range mapp {
		if a == len {
			return true
		}
	}

	return false

}
