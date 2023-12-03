package main

import "testing"

func Test_linearSearch(t *testing.T) {
	type args struct {
		arr []int
		key int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Element found in the array",
			args: args{arr: []int{1, 2, 3, 4, 5}, key: 3},
			want: true,
		},
		{
			name: "Element not found in the array",
			args: args{arr: []int{1, 2, 3, 4, 5}, key: 6},
			want: false,
		},
		{
			name: "Empty array",
			args: args{arr: []int{}, key: 3},
			want: false,
		},
		{
			name: "Element found in an array with repeated elements",
			args: args{arr: []int{1, 2, 2, 3, 3, 4, 5}, key: 3},
			want: true,
		},
		{
			name: "Element not found in an array with repeated elements",
			args: args{arr: []int{1, 2, 2, 3, 3, 4, 5}, key: 6},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := linearSearch(tt.args.arr, tt.args.key); got != tt.want {
				t.Errorf("linearSearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkLinearSearch measures the performance of the linearSearch function.
func BenchmarkLinearSearch(b *testing.B) {
	// Create a large array with a known element to search for
	arr := make([]int, 1000000)
	for i := 0; i < 1000000; i++ {
		arr[i] = i + 1
	}
	key := len(arr) / 2

	// Run the benchmark function b.N times
	for i := 0; i < b.N; i++ {
		// Call the linearSearch function with the target element
		linearSearch(arr, key)
	}
}

/*

mayurwadekar@Mayurs-MacBook-Pro binarysearch % go test -bench .
goos: darwin
goarch: arm64
pkg: Go_Exercise/SearchingAlgorithms/binarysearch
BenchmarkBinarySearch-8         42458973                27.97 ns/op
PASS
ok      Go_Exercise/SearchingAlgorithms/binarysearch    2.278s

*/
