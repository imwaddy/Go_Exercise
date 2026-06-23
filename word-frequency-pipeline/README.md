# Exercise: Word Frequency Pipeline

**Difficulty:** Beginner-Intermediate  
**Concepts:** goroutines, channels, pipeline pattern, fan-out, aggregation

---

## Problem Statement

You are given a list of sentences. Count the frequency of each word across all sentences concurrently.

Pipeline:

```
sentences → split into words → count frequency → print top 5 words
```

- Each sentence is processed by a separate goroutine
- Words flow through channels
- Final aggregation collects all words and counts them
- Print top 5 most frequent words sorted by count descending
- Tie in count → sort alphabetically ascending

---

## Input

```go
sentences := []string{
    "the cat sat on the mat",
    "the cat sat on the hat",
    "the cat in the hat",
    "the fat cat sat",
    "cat cat cat hat",
}
```

---

## Sample Output

```
cat 8
the 8
sat 4
hat 3
on 2
```

---

## Constraints

- No mutexes — use only channels for communication
- Each sentence processed concurrently
- Words are all lowercase, no punctuation

---

## Things to figure out

- How many channels do you need?
- How do you know when all goroutines are done?
- How do you safely aggregate without a mutex?

---

## Follow-Up

1. Add a `filterStopWords` stage — drop words like "the", "on", "in"
2. Print bottom 5 instead of top 5
3. Handle punctuation — strip `,` `.` `!` from words

---

## Verification

```bash
go run main.go
go test -race ./...
```
