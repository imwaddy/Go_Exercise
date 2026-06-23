# Exercise: Number Pipeline

**Difficulty:** Beginner  
**Concepts:** goroutines, channels, pipeline pattern, channel chaining

---

## Problem Statement

Build a 3-stage pipeline using goroutines and channels:

```
generate → filter → square → print
```

- **Stage 1 — generate:** send numbers 1 to N into a channel
- **Stage 2 — filter:** pass only even numbers
- **Stage 3 — square:** square each number, send to output

Main reads from final channel and prints results.

---

## Requirements

1. Each stage is a separate goroutine
2. Each stage takes an input channel, returns an output channel
3. No slices — values flow one at a time through channels
4. Stages run concurrently — stage 2 starts processing before stage 1 finishes
5. Pipeline shuts down cleanly when generator is done

---

## Function Signatures

```go
func generate(n int) <-chan int
func filterEven(in <-chan int) <-chan int
func square(in <-chan int) <-chan int
```

---

## Sample Input

```
N = 10
```

## Sample Output

```
4
16
36
64
100
```

(evens: 2,4,6,8,10 → squared: 4,16,36,64,100)

---

## How stages connect

```go
nums    := generate(10)
evens   := filterEven(nums)
squared := square(evens)

for v := range squared {
    fmt.Println(v)
}
```

---

## Things to figure out

- When does each channel get closed?
- Who closes which channel?
- What happens if you forget `close()`?

---

## Follow-Up

1. Add a 4th stage: `sum` — collect all values and print total
2. Add context cancellation — cancel pipeline midway
3. Fan-out: run 3 parallel `square` goroutines instead of one

---

## Verification

```bash
go run main.go
go test -race ./...
```
