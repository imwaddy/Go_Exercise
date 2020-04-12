// Author: Mayur Wadekar

package main

import (
	"testing"
)

func Test_isPalindrome(t *testing.T) {
	type args struct {
		stringToCheck string
	}
	tests := []struct {
		name string
		args args
	}{
		{"Palindrome string", args{stringToCheck: "nitin"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			isPalindrome(tt.args.stringToCheck)
		})
	}
}

func Test_isPalindromeOne(t *testing.T) {
	type args struct {
		stringToCheck string
	}
	tests := []struct {
		name string
		args args
	}{
		{"Test_isPalindromeOne", args{stringToCheck: "a"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			isPalindrome(tt.args.stringToCheck)
		})
	}
}

func Test_isPalindromeTwo(t *testing.T) {
	type args struct {
		stringToCheck string
	}
	tests := []struct {
		name string
		args args
	}{
		{"Test_isPalindromeTwo", args{stringToCheck: "ab"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			isPalindrome(tt.args.stringToCheck)
		})
	}
}

func Test_isPalindromeThree(t *testing.T) {
	type args struct {
		stringToCheck string
	}
	tests := []struct {
		name string
		args args
	}{
		{"Test_isPalindromeThree", args{stringToCheck: "bb"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			isPalindrome(tt.args.stringToCheck)
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"main function"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
