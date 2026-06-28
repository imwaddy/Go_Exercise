package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type Result struct {
	worker int
	number int
}

func main() {
	jobs := make(chan int, 4)
	results := make(chan Result)

	var wg sync.WaitGroup

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	go func() {
		defer close(jobs)
		for i := 0; ; i++ {
			select {
			case <-ctx.Done():
				return
			case jobs <- i:
			}
		}
	}()

	for j := 0; j < 6; j++ {
		wg.Add(1)
		go worker(ctx, j, jobs, results, &wg)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for r := range results {
		fmt.Println("Result worker", r.worker, "prints", r.number)
	}

	fmt.Println("Main exiting")
}

func worker(ctx context.Context, wid int, jobs <-chan int, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			return

		case job, ok := <-jobs:
			if !ok {
				return
			}

			select {
			case <-ctx.Done():
				return
			case results <- Result{worker: wid, number: job}:
			}
		}
	}
}
