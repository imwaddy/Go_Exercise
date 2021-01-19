// You have to read a file and each row of file contains the number. We have to read a file concurrently and have to append the sum of numbers into the file. If first file contains the filename then also calculate the sum and add with the first file ex.

// file.txt
// 10
// 20
// file2.txt
// 30

// file2.txt
// 10

// output
// file.txt
// 10
// 20
// 30
// 70

// Sum of file.txt is 60 and file2.txt is 10 total 70 should write into file.txt

package main

import (
	"os"
	"reflect"
	"sync"
	"testing"
)

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "Test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

func Test_readFilePositive(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name     string
		args     args
		wantFile *os.File
	}{
		{name: "Test readFilePositive", args: args{
			filename: "file.txt",
		},
			wantFile: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotFile := readFile(tt.args.filename); reflect.DeepEqual(gotFile, tt.wantFile) {
				t.Errorf("readFile() = %v, want %v", gotFile, tt.wantFile)
			}
		})
	}
}

func Test_readFileNegative(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name     string
		args     args
		wantFile *os.File
	}{
		{name: "Test readFileNegative", args: args{
			filename: "file.txt",
		},
			wantFile: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotFile := readFile(tt.args.filename); !reflect.DeepEqual(gotFile, tt.wantFile) {
				t.Errorf("readFile() = %v, want %v", gotFile, tt.wantFile)
			}
		})
	}
}

func Test_readTextPositive(t *testing.T) {

	file, _ := os.Open("file.txt")
	var wg sync.WaitGroup
	var mutex sync.Mutex
	type args struct {
		file  *os.File
		wg    *sync.WaitGroup
		mutex *sync.Mutex
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "Test readTextPositive", args: args{
			file:  file,
			wg:    &wg,
			mutex: &mutex,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			readText(tt.args.file, tt.args.wg, tt.args.mutex)
		})
	}
}
