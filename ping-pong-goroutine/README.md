# Exercise 1: Ping-Pong Goroutines

**Difficulty:** Beginner  
**Concepts:** goroutines, bidirectional channels, channel direction, synchronization

---

## Problem Statement

Two goroutines pass a counter back and forth through channels.

- Goroutine **A** sends "ping" and increments counter
- Goroutine **B** receives ping, sends "pong" back
- Repeat for exactly **N** rounds
- After N rounds, both goroutines exit cleanly
- Main prints each exchange in order

---

## Requirements

1. Use **two unidirectional channels**: one for ping direction, one for pong direction
2. No global variables
3. No `time.Sleep` for synchronization
4. Must handle N=0 (no output, no deadlock)

---

## Function Signatures

```go
func pinger(ping chan<- string, pong <-chan string, n int)
func ponger(ping <-chan string, pong chan<- string, n int)
```

---

## Sample Input

```
N = 3
```

## Sample Output

```
Round 1: ping -> pong
Round 2: ping -> pong
Round 3: ping -> pong
done
```

---

## Constraints

```
0 <= N <= 1,000,000
```

---

## Follow-Up

- Add a third goroutine **C** that can "intercept" every 3rd ping and reply with "intercept" instead of "pong"
- Use a `done` channel instead of a counter to stop goroutines

---

## Verification

```bash
go run main.go
go test -race ./...
```
