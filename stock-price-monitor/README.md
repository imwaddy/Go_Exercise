# Exercise: Stock Price Monitor

**Difficulty:** Beginner-Intermediate
**Concepts:** goroutines, channels, fan-in, select, done channel

---

## Problem Statement

You have 3 stocks. Each stock has a goroutine that emits price updates every few milliseconds.

Collect all price updates from all 3 stocks into a single stream and print them as they arrive.

Stop collecting after **N total updates** received across all stocks.

---

## Input

```go
stocks := []struct {
    name   string
    prices []int
}{
    {"AAPL", []int{150, 152, 149, 155, 160}},
    {"GOOG", []int{2800, 2750, 2820, 2795, 2810}},
    {"TSLA", []int{700, 710, 695, 720, 715}},
}

N := 10  // stop after 10 updates
```

---

## Sample Output (order will vary — concurrent)

```
AAPL: 150
GOOG: 2800
TSLA: 700
AAPL: 152
TSLA: 710
GOOG: 2750
AAPL: 149
GOOG: 2820
TSLA: 695
AAPL: 155
received 10 updates, stopping
```

---

## Requirements

1. Each stock runs in its own goroutine
2. All updates flow into one merged channel (fan-in)
3. Use a `done` channel to signal goroutines to stop
4. After N updates — close `done`, stop all goroutines, exit cleanly
5. No goroutine leak after exit

---

## Things to figure out

- How does each stock goroutine know when to stop?
- How do you merge 3 channels into 1?
- How do you count N updates and stop cleanly?

---

## Constraints

```
1 <= N <= total prices across all stocks
Each stock goroutine must respect done channel
```

---

## Follow-Up

1. Add `time.Sleep(random ms)` in each stock goroutine to simulate real feed
2. Instead of stopping after N updates, stop after a timeout (use `time.After`)
3. Track latest price per stock — print summary at the end

---

## Verification

```bash
go run main.go
go test -race ./...
```
