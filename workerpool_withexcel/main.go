package main

import (
	"fmt"
	"sort"
	"strings"
	"sync"

	excelize "github.com/xuri/excelize/v2"
)

var fileName = "Book-1.xlsx"

type workerResult struct {
	id   int
	rows [][]string
}

func main() {
	f, err := excelize.OpenFile(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}

	const numJobs = 2
	jobs := make(chan [][]string, numJobs)
	results := make(chan workerResult, numJobs)

	var wg sync.WaitGroup
	for w := 1; w <= numJobs; w++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			worker(id, jobs, results)
		}(w)
	}

	chunkSize := len(rows) / numJobs
	for i := 0; i < numJobs; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if i == numJobs-1 {
			end = len(rows)
		}
		jobs <- rows[start:end]
	}
	close(jobs)

	go func() {
		wg.Wait()
		close(results)
	}()

	collected := make([]workerResult, 0, numJobs)
	for result := range results {
		collected = append(collected, result)
	}

	sort.Slice(collected, func(i, j int) bool {
		return collected[i].id < collected[j].id
	})

	for i, result := range collected {
		if i > 0 {
			fmt.Println()
		}
		process(result)
	}
}

func worker(id int, jobs <-chan [][]string, results chan<- workerResult) {
	for input := range jobs {
		fmt.Printf("Worker %d processing %d rows\n", id, len(input))
		results <- workerResult{id: id, rows: input}
	}
}

func process(r workerResult) {
	fmt.Printf("Worker %d:\n", r.id)
	for _, row := range r.rows {
		fmt.Println(" ", strings.Join(row, "\t"))
	}
}
