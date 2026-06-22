package main

import (
	"fmt"
	"sync"
)

func main() {
	toOdd := make(chan struct{})
	toEven := make(chan struct{})
	var wg sync.WaitGroup
	N := 10
	wg.Add(2)

	go oddWorker(toOdd, toEven, N, &wg)
	go evenWorker(toOdd, toEven, N, &wg)

	toOdd <- struct{}{}
	wg.Wait()
}

func oddWorker(toOdd <-chan struct{}, toEven chan<- struct{}, n int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= n; i = i + 2 {
		<-toOdd
		fmt.Println("Number:", i)
		toEven <- struct{}{}
	}
}
func evenWorker(toOdd chan<- struct{}, toEven <-chan struct{}, n int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 2; i <= n; i = i + 2 {
		<-toEven
		fmt.Println("Number:", i)
		if i < n {
			toOdd <- struct{}{}
		}

	}
}
