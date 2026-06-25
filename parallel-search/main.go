package main

import (
	"fmt"
	"os"
	"strings"
	"sync"
)

type Result struct {
	WorkerID int
	fileName string
}

func main() {
	files := []string{
		"a.txt", "b.txt", "c.txt", "d.txt", "e.txt",
	}

	jobs := make(chan string, 2)
	done := make(chan struct{})
	results := make(chan Result)
	var wg sync.WaitGroup

	keyword := "error"
	workers := 3
	maxResults := 2

	dispatcher(files, jobs, done)

	for i := 1; i <= workers; i++ {
		wg.Add(1)
		go worker(i, jobs, results, keyword, done, &wg)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for r := range results {
		fmt.Printf("[Worker %d] match %s\n", r.WorkerID, r.fileName)
		maxResults--
		if maxResults == 0 {
			close(done)
			break
		}
	}
}

func worker(id int, jobs <-chan string, results chan<- Result, keyword string, done <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		if search(j, keyword) {
			results <- Result{fileName: j, WorkerID: id}
		}
	}
}

func search(filename, keyword string) bool {
	f, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error opening file ", err)
		return false
	}

	return strings.Contains(string(f), keyword)
}

func dispatcher(files []string, jobs chan<- string, done <-chan struct{}) {
	go func() {
		for _, f := range files {
			select {
			case <-done:
				return
			case jobs <- f:
			}
		}
		close(jobs)
	}()
}
