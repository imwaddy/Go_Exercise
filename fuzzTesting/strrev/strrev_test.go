package main

import (
	"strings"
	"testing"
)

func FuzzStrRev(f *testing.F) {

	f.Fuzz(func(t *testing.T, str string) {

		reverseString("俺夜嵐")

		s := reverseString(str)
		rev := reverse(str)

		if s != rev {
			t.Errorf("Val s:[%s] and Val rev: [%s]", s, rev)
			t.Fail()
		}

	})

}

func reverse(str string) string {
	l := len(str)
	result := make([]string, l)

	for _, c := range str {
		l--
		result[l] = string(c)
	}
	return strings.Join(result, "")
}
