# Exercise 1: Graceful Shutdown Worker Pool

**Difficulty:** Beginner–Intermediate  
**Concepts:** goroutines, channels, `context`, `os/signal`, graceful shutdown, `sync.WaitGroup`

---

## Problem Statement

Write a Go program that starts:

- one **job producer**
- multiple **worker goroutines**
- one **result printer**

The producer continuously generates jobs and sends them to a jobs channel.  
Each worker receives jobs, processes them, and sends results to a results channel.

When the user presses **`Ctrl+C`** in the terminal/command prompt:

- the program should stop accepting new jobs
- workers should finish promptly and exit cleanly
- the results channel should be closed only after all workers are done
- the main goroutine should print all completed results and terminate without leaking goroutines

---

## Requirements

1. Use `os/signal` to catch `Ctrl+C` (`os.Interrupt`) and `SIGTERM`
2. Use `context.Context` to broadcast shutdown to all goroutines
3. Use a `sync.WaitGroup` to wait for all workers to finish
4. Use channels for job distribution and result collection
5. Do **not** use `time.Sleep` for synchronization
6. The program must terminate cleanly when interrupted, even if workers are idle or the jobs channel is full

---

## Behavior

- The producer keeps sending incrementing integers as jobs
- Each worker receives a job and produces a result containing:
  - the worker id
  - the job number
- The main goroutine prints results as they arrive
- On `Ctrl+C`, all goroutines should stop gracefully

---

## Function Signatures

```go
type Result struct {
	worker int
	number int
}

func worker(ctx context.Context, wid int, jobs <-chan int, results chan<- Result, wg *sync.WaitGroup)
```
---

## Sample Output

```text
Result worker 2 prints 0
Result worker 0 prints 1
Result worker 4 prints 2
Result worker 1 prints 3
Shutdown signal received
Worker 3 stopping
Worker 5 stopping
Main exiting
```
---

> Note: Output order is not guaranteed because goroutines run concurrently.

---

## Constraints

```text
1 <= number of workers <= 100
0 <= jobs generated until shutdown
```

---

## Rules

1. Workers must stop when:
   - the context is canceled, or
   - the jobs channel is closed
2. The producer must close the jobs channel when it stops sending
3. The results channel must be closed only after all workers finish
4. The main goroutine must range over the results channel until it is closed

---

## Common Pitfalls

- Closing the results channel too early
- Reading from the results channel inside a worker
- Calling `wg.Done()` at the start instead of using `defer wg.Done()`
- Not selecting on `ctx.Done()` while sending to or receiving from channels
- Forgetting to close `jobs`
- Using the wrong worker id when starting goroutines

---

## Follow-Up

- Add a bounded job queue with backpressure
- Make workers simulate processing delays
- Support a `done` channel instead of `context.Context`
- Add metrics for:
  - number of jobs produced
  - number of jobs processed
  - number of workers stopped by shutdown

---

## Verification

```bash
go run main.go
go test -race ./...
```

---

## Stretch Goal

Modify the solution so that:

- workers stop immediately on `Ctrl+C`
- the producer stops even if the jobs channel buffer is full
- no goroutine is left blocked on send or receive

---



