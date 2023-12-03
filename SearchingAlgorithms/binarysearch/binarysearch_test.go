package main

import "testing"

func TestBinarySearch(t *testing.T) {
	type args struct {
		arr []int
		e   int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Element found in a sorted array",
			args: args{arr: []int{1, 2, 3, 4, 5}, e: 3},
			want: true,
		},
		{
			name: "Element not found in a sorted array",
			args: args{arr: []int{1, 2, 3, 4, 5}, e: 6},
			want: false,
		},
		{
			name: "Empty array",
			args: args{arr: []int{}, e: 3},
			want: false,
		},
		{
			name: "Element found in an array with repeated elements",
			args: args{arr: []int{1, 2, 2, 3, 3, 4, 5}, e: 3},
			want: true,
		},
		{
			name: "Element not found in an array with repeated elements",
			args: args{arr: []int{1, 2, 2, 3, 3, 4, 5}, e: 6},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinarySearch(tt.args.arr, tt.args.e); got != tt.want {
				t.Errorf("BinarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkBinarySearch measures the performance of the BinarySearch function.
func BenchmarkBinarySearch(b *testing.B) {
	// Create a sorted array with a large number of elements
	arr := make([]int, 1000000)
	for i := 0; i < 1000000; i++ {
		arr[i] = i + 1
	}

	// Run the benchmark function b.N times
	for i := 0; i < b.N; i++ {
		// Call the BinarySearch function with the target element (middle element)
		BinarySearch(arr, len(arr)/2)
	}
}

/*

mayurwadekar@Mayurs-MacBook-Pro binarysearch % go test -bench .
goos: darwin
goarch: arm64
pkg: Go_Exercise/SearchingAlgorithms/binarysearch
BenchmarkBinarySearch-8         42496522                28.03 ns/op
PASS
ok      Go_Exercise/SearchingAlgorithms/binarysearch    2.201s

*/
