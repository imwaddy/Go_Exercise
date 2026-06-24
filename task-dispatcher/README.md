# Exercise: Task Dispatcher

**Difficulty:** Intermediate
**Concepts:** worker pool, buffered channels, WaitGroup, result collection

---

## Problem Statement

You have a list of tasks. Each task has an ID and a duration (ms).
Dispatch tasks to a fixed pool of **W workers**.
Each worker simulates work with `time.Sleep(duration)`.
Collect all results and print them sorted by task ID.

---

## Input

```go
tasks := []Task{
    {ID: 1, Duration: 80},
    {ID: 2, Duration: 20},
    {ID: 3, Duration: 50},
    {ID: 4, Duration: 10},
    {ID: 5, Duration: 60},
    {ID: 6, Duration: 30},
    {ID: 7, Duration: 90},
    {ID: 8, Duration: 40},
}

W := 3  // worker pool size
```

---

## Sample Output (sorted by ID)

```
task 1 done in 80ms by worker 1
task 2 done in 20ms by worker 2
task 3 done in 50ms by worker 3
task 4 done in 10ms by worker 2
task 5 done in 60ms by worker 2
task 6 done in 30ms by worker 3
task 7 done in 90ms by worker 1
task 8 done in 40ms by worker 3
```

---

## Requirements

1. Fixed pool of W workers — not one goroutine per task
2. Tasks dispatched via a **buffered** channel
3. Results collected via a results channel
4. Final output sorted by task ID
5. Main waits for ALL tasks to complete before printing

---

## Types

```go
type Task struct {
    ID       int
    Duration int // ms
}

type Result struct {
    TaskID   int
    Duration int
    WorkerID int
}
```

---

## Things to figure out

- Who closes the tasks channel and when?
- How does main know all workers are done?
- How do you collect results without knowing order of completion?

---

## Constraints

```
1 <= W <= 10
1 <= len(tasks) <= 1000
1 <= Duration <= 200
```

---

## Follow-Up

1. Add context cancellation — cancel all pending tasks after timeout
2. Track which worker processed most tasks
3. Make buffer size configurable — observe how it affects throughput

---

## Verification

```bash
go run main.go
go test -race ./...
```
