# Exercise: Odd-Even Sequencer

**Difficulty:** Beginner-Intermediate  
**Concepts:** goroutines, unbuffered/buffered channels, token passing, sequential coordination

---

## Problem Statement

Print numbers from **1 to N** in order using exactly **two goroutines**:

- Goroutine **A** prints only odd numbers: 1, 3, 5, ...
- Goroutine **B** prints only even numbers: 2, 4, 6, ...
- Output must be strictly sequential: `1 2 3 4 5 ...`
- No `time.Sleep` for ordering
- No shared variables
- Coordination only via channels

---

## Requirements

1. Exactly 2 goroutines — one odd, one even
2. Use channels as **turn tokens** (not for passing the numbers themselves)
3. Main goroutine kicks off the sequence then waits
4. Both goroutines exit cleanly after N numbers printed
5. No goroutine leak

---

## Function Signatures

```go
func oddWorker(toOdd <-chan struct{}, toEven chan<- struct{}, n int, wg *sync.WaitGroup)
func evenWorker(toOdd chan<- struct{}, toEven <-chan struct{}, n int, wg *sync.WaitGroup)
```

---

## Sample Input

```
N = 10
```

## Sample Output

```
1
2
3
4
5
6
7
8
9
10
```

---

## Constraints

```
1 <= N <= 1,000,000
N is always even
```

---

## Things to figure out

- Which channel needs a buffer and why?
- Who sends the first token and when?
- How does the last goroutine exit without sending a token that no one reads?

---

## Follow-Up

1. Extend to **3 goroutines**: divisible by 3, remainder 1, remainder 2 — print 1 to N in order
2. Make N odd — how does your exit condition change?
3. Replace `sync.WaitGroup` with a `done` channel

---

## Verification

```bash
go run main.go
go test -race ./...
```
