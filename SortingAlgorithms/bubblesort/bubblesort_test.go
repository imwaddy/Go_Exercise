package main

import (
	"reflect"
	"testing"
)

func Test_bubbleSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int // The expected sorted array
	}{
		{
			name: "Unsorted array",
			args: args{arr: []int{9, 5, 4, 3, 21, 3}},
			want: []int{3, 3, 4, 5, 9, 21},
		},
		{
			name: "Sorted array",
			args: args{arr: []int{1, 2, 3, 4, 5}},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "Reversed array",
			args: args{arr: []int{5, 4, 3, 2, 1}},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "Array with repeated elements",
			args: args{arr: []int{3, 2, 1, 3, 4, 2, 5, 1}},
			want: []int{1, 1, 2, 2, 3, 3, 4, 5},
		},
		{
			name: "Empty array",
			args: args{arr: []int{}},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bubbleSort(tt.args.arr)
			if !reflect.DeepEqual(tt.args.arr, tt.want) {
				t.Errorf("bubbleSort() = %v, want %v", tt.args.arr, tt.want)
			}
		})
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	// Create a sample input for the benchmark
	inputArray := []int{9, 5, 4, 3, 21, 3}

	// Run the benchmark function b.N times
	for i := 0; i < b.N; i++ {
		// Call the bubbleSort function with the sample input
		bubbleSort(inputArray)
	}
}

/*

mayurwadekar@Mayurs-MacBook-Pro bubblesort % go test -bench .
goos: darwin
goarch: arm64
pkg: Go_Exercise/SortingAlgorithms/bubblesort
BenchmarkBubbleSort-8           97074414                12.22 ns/op
PASS
ok      Go_Exercise/SortingAlgorithms/bubblesort        2.062s

*/
