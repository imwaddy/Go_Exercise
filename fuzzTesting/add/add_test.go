package main

import (
	"testing"
)

func FuzzAdd(f *testing.F) {
	f.Fuzz(func(t *testing.T, no1, no2 int) {
		add(19, 20)
		res := add(no1, no2)

		if res >= 400 {
			t.Errorf("%v", res)
		}

	})
}
