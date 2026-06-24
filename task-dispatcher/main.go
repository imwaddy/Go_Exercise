package main

import (
	"context"
	"fmt"
	"sort"
	"sync"
	"time"
)

type Task struct {
	ID       int `json:"ID"`
	Duration int `json:"Duration"`
}

type Result struct {
	Task
	WorkerID int
}

func main() {

	tasks := []Task{
		{ID: 3, Duration: 50},
		{ID: 8, Duration: 40},
		{ID: 1, Duration: 80},
		{ID: 2, Duration: 20},
		{ID: 4, Duration: 10},
		{ID: 5, Duration: 60},
		{ID: 6, Duration: 30},
		{ID: 7, Duration: 90},
	}

	W := 3
	var wg sync.WaitGroup

	jobs := make(chan Task, 3)
	result := make(chan Result, len(tasks))

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(30*time.Second))
	defer cancel()

	go func() {
		for _, task := range tasks {
			jobs <- task
		}
		close(jobs)
	}()

	for i := 0; i < W; i++ {
		wg.Add(1)
		go worker(ctx, i, jobs, result, &wg)
	}

	go func() {
		wg.Wait()
		close(result)
	}()

	data := make([]Result, 0, len(tasks))
	for r := range result {
		data = append(data, r)
	}

	sort.Slice(data, func(i, j int) bool {
		return data[i].ID < data[j].ID
	})

	for _, d := range data {
		fmt.Println("Task ", d.ID, " done in ", d.Duration, "ms by worker ", d.WorkerID)
	}
}

func worker(ctx context.Context, workerID int, jobs chan Task, result chan Result, wg *sync.WaitGroup) {
	defer wg.Done()
	select {
	case <-ctx.Done():
		return
	case job, ok := <-jobs:
		if !ok {
			return
		}
		time.Sleep(time.Duration(job.Duration) * time.Millisecond)
		result <- Result{Task: job, WorkerID: workerID}
	}
}
