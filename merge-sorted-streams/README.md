# Exercise 2: Merge Sorted Streams

**Difficulty:** Intermediate  
**Concepts:** fan-in pattern, goroutines, channel merge, ordering, done channel

---

## Problem Statement

You have **K sorted streams** of integers (each stream is a goroutine sending values on a channel).  
Merge all streams into a **single sorted output stream**.

---

## Requirements

1. Each stream goroutine sends integers in ascending order then closes its channel
2. Write a `merge(streams ...<-chan int) <-chan int` function
3. Output channel must emit all integers in ascending order
4. Use goroutines and channels — no buffering the entire stream into a slice first
5. Support cancellation via `context.Context` — if context cancelled, all goroutines exit cleanly

---

## Function Signatures

```go
func streamGenerator(nums []int) <-chan int
func merge(ctx context.Context, streams ...<-chan int) <-chan int
```

---

## Sample Input

```
Stream 1: [1, 4, 7, 10]
Stream 2: [2, 5, 8]
Stream 3: [3, 6, 9, 12]
```

## Sample Output

```
1 2 3 4 5 6 7 8 9 10 12
```

---

## Hint on approach

Think about how to pick the minimum from K channel heads at each step.  
A min-heap over channel values is one approach.

---

## Constraints

```
1 <= K <= 100
1 <= len(each stream) <= 100,000
values fit in int
```

---

## Follow-Up

- What if streams are NOT pre-sorted? Modify to sort within each stream goroutine first.
- Add a `done` channel instead of context for cancellation.

---

## Verification

```bash
go run main.go
go test -race ./...
```
