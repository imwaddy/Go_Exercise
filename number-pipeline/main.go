package main

import "fmt"

func main() {
	N := 10

	generated := generate(N)
	filter := filterEven(generated)
	square := square(filter)

	for r := range square {
		fmt.Println(r)
	}

}

func generate(n int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 1; i <= n; i++ {
			ch <- i
		}
	}()
	return ch
}

func filterEven(in <-chan int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := range in {
			if i%2 == 0 {
				ch <- i
			}
		}
	}()
	return ch
}

func square(in <-chan int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := range in {
			ch <- i * i
		}
	}()
	return ch
}
