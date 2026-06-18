package main

import (
	"fmt"
	"sync"
	"time"

	excelize "github.com/xuri/excelize/v2"
)

var mutex *sync.Mutex
var rowsread = -1
var fileName = "Book-1.xlsx"
var start, end int

func main() {

	// openFile(fileName)

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
	jobs := make(chan int, numJobs)
	results := make(chan [][]string, numJobs)

	r := len(rows) / numJobs
	start = 1
	end = r + 1

	for w := 1; w <= numJobs; w++ {
		go worker(w, jobs, results, rows[start:end])
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
		fmt.Println("jobs ", <-results)
		// process(<-results)
	}
	close(jobs)

	for a := 1; a <= numJobs; a++ {
		// fmt.Println(<-results)
		process(<-results)
	}

	time.Sleep(5 * time.Second)

}

func worker(id int, jobs <-chan int, results chan<- [][]string, input [][]string) {
	fmt.Println("input ", input)
	for w, _ := range jobs {
		if w != 1 {
			start = end
			end += r
		}
		fmt.Println("Start : ", start)
		fmt.Println("End : ", end)
		work(results, input)
	}
}

func work(results chan<- [][]string, input [][]string) {
	// getExcelRows()
	results <- input
}

func process(s [][]string) {
	fmt.Println(s)
}
