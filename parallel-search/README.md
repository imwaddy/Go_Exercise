# Exercise: Parallel File Search

**Difficulty:** Intermediate  
**Concepts:** fan-out, worker pool, buffered channels, `select`, done channel, early exit

---

## Problem Statement

You have a list of filenames and a search keyword.  
Search all files **concurrently** and return every file that contains the keyword.

```
files ──► dispatcher ──► [worker][worker][worker] ──► results
```

---

## Requirements

1. Fixed pool of **W workers** (not one goroutine per file)
2. Dispatcher sends filenames into a `jobs` channel
3. Each worker reads from `jobs`, searches file content, sends match to `results`
4. Main collects all matches and prints them
5. All workers exit cleanly when jobs are exhausted — no goroutine leak
6. Support **early exit**: if `maxResults` is hit, cancel remaining work via done channel

---

## Function Signatures

```go
func dispatcher(files []string, jobs chan<- string, done <-chan struct{})
func worker(id int, jobs <-chan string, results chan<- string, keyword string, done <-chan struct{}, wg *sync.WaitGroup)
func search(filename, keyword string) bool  // returns true if file contains keyword
```

---

## Sample Input

```go
files := []string{
    "a.txt", "b.txt", "c.txt", "d.txt", "e.txt",
}
keyword    := "error"
workers    := 3
maxResults := 2   // stop after finding 2 matches
```

## Sample Output

```
[worker-1] match: a.txt
[worker-3] match: c.txt
early exit: max results reached
```

---

## Things to figure out

- Who closes `jobs` channel — dispatcher or main?
- What buffered size makes sense for `jobs`? Why?
- How does `done` channel signal both dispatcher AND workers simultaneously?
- What happens if you close `done` but a worker is mid-send to `results`?

---

## Progression from Part 1

| Part 1 (chat-room-hub) | Part 2 (parallel-search) |
|------------------------|--------------------------|
| fan-in: many → one | fan-out: one → many |
| WaitGroup for senders | WaitGroup for receivers |
| unbuffered channel | buffered jobs channel |
| no cancellation | done channel + early exit |

---

## Constraints

```
1 <= W <= 10      workers
1 <= F <= 10,000  files
```

---

## Verification

```bash
go run main.go
go test -race ./...
```

### Race detector mandatory — must pass clean.

---

## Follow-Up

- Replace done channel with `context.WithCancel`
- Add per-worker stats: files scanned, matches found
- Make `maxResults = 0` mean "no limit"
