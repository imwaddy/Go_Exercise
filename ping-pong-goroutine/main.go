package main

import (
	"fmt"
	"sync"
)

func main() {
	var rounds int
	fmt.Printf("Enter number of rounds:")
	fmt.Scanf("%d", &rounds)

	var ping = make(chan string)
	var pong = make(chan string)
	var wg sync.WaitGroup
	wg.Add(2)

	go pinger(ping, pong, rounds, &wg)
	go ponger(ping, pong, rounds, &wg)

	wg.Wait()
}

func pinger(ping chan<- string, pong <-chan string, n int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < n; i++ {
		ping <- "ping"
		<-pong
		fmt.Printf("Round %d: ping -> pong\n", i+1)
	}
}

func ponger(ping <-chan string, pong chan<- string, n int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < n; i++ {
		<-ping
		pong <- "pong"
	}
}
