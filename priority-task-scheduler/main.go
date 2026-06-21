package main

import (
	"container/heap"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

type Task struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Priority   int    `json:"priority"`
	DurationMs int    `json:"duration_ms"`
}

type Result struct {
	Task      Task
	WorkerID  int
	Status    string // "completed" | "cancelled"
	StartedAt time.Time
	EndedAt   time.Time
}

type PriorityQueue []*Task

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Priority < pq[j].Priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x any) {
	item := x.(*Task)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	h := *pq
	n := len(h)
	item := h[n-1]
	h[n-1] = nil
	*pq = h[:n-1]
	return item
}

func LoadTasks(path string) ([]Task, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		log.Printf("Error while reading file %+v", err)
		return []Task{}, err
	}

	var tasks []Task
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		log.Printf("Error while unmarshaling %+v", err)
		return []Task{}, err
	}
	return tasks, err
}

func main() {
	tasks, err := LoadTasks("input/tasks.json")
	if err != nil {
		return
	}

	noOfWorkers := 5

	pq := make(PriorityQueue, 0, len(tasks))
	heap.Init(&pq)

	for i := range tasks {
		heap.Push(&pq, &tasks[i])
	}

	result := make(chan Result, len(tasks))
	jobs := make(chan Task, 3)

	ctx, cancel := context.WithTimeout(context.TODO(), time.Duration(30*time.Second))
	defer cancel()

	start := time.Now()

	for i := 0; i < noOfWorkers; i++ {
		go worker(ctx, jobs, result, i)
	}

	go func() {
		for pq.Len() > 0 {
			jobs <- *heap.Pop(&pq).(*Task)
		}
		close(jobs)
	}()

	var results []Result
	for i := 0; i < len(tasks); i++ {
		r := <-result
		fmt.Printf("[worker-%d] %-20s priority=%-2d status=%s duration=%dms\n",
			r.WorkerID, r.Task.Name, r.Task.Priority, r.Status,
			r.EndedAt.Sub(r.StartedAt).Milliseconds())
		results = append(results, r)
	}

	PrintSummary(results, time.Since(start), noOfWorkers)

}

func PrintSummary(results []Result, wallTime time.Duration, noOfWorkers int) {
	completed, cancelled := 0, 0
	for _, r := range results {
		if r.Status == "completed" {
			completed++
		} else {
			cancelled++
		}
	}

	fmt.Println("\n=== Summary ===")
	fmt.Printf("  Total tasks    : %d\n", len(results))
	fmt.Printf("  Completed      : %d\n", completed)
	fmt.Printf("  Failed/Timeout : %d\n", cancelled)
	fmt.Printf("  Total wall time: ~%dms (%d workers in parallel)\n", wallTime.Milliseconds(), noOfWorkers)
}

func worker(ctx context.Context, jobs chan Task, result chan Result, workerID int) {
	for job := range jobs {
		startedAt := time.Now().UTC()
		select {
		case <-ctx.Done():
			result <- Result{
				Task:      job,
				WorkerID:  workerID,
				Status:    "cancelled",
				StartedAt: time.Now().UTC(),
				EndedAt:   time.Now().UTC(),
			}
			return
		default:
			time.Sleep(time.Duration(job.DurationMs) * time.Millisecond)
			result <- Result{
				Task:      job,
				WorkerID:  workerID,
				Status:    "completed",
				StartedAt: startedAt,
				EndedAt:   time.Now().UTC(),
			}
		}
	}

}
