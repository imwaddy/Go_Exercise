# Problem: Multi-Stage Pipeline Processor

## Problem Statement

Build a concurrent multi-stage data pipeline in Go that reads records from a JSON file, passes them through a chain of processing stages, and writes results to stdout. Each stage runs in its own goroutine pool. Backpressure flows naturally via buffered channels.

---

## Requirements

### Core

1. **Read** records from `input/records.json`.
2. **Three pipeline stages** (each with its own worker pool):
   - **Stage 1 — Validate**: Drop records with missing/invalid fields. Pass valid ones forward.
   - **Stage 2 — Enrich**: Add a computed field `score` (see formula below). Simulate 50ms work.
   - **Stage 3 — Aggregate**: Group enriched records by `category`. Simulate 30ms work.
3. **Worker counts** — configurable per stage (defaults: validate=2, enrich=3, aggregate=2).
4. **Context cancellation** — global timeout (default: 5s). All stages drain and exit cleanly on cancel.
5. **Result** — print per-category aggregates when pipeline finishes.

### Score Formula (Stage 2)

```
score = (amount * 0.4) + (quantity * 0.6)
```

### Validation Rules (Stage 1 — drop record if ANY fail)

- `id` must be non-empty string
- `amount` must be > 0
- `quantity` must be >= 1
- `category` must be one of: `"electronics"`, `"clothing"`, `"food"`, `"books"`

### Aggregation (Stage 3)

Per category, accumulate:
- `count` — number of records
- `total_score` — sum of scores
- `avg_score` — total_score / count

### Constraints

- Stages connected via channels — no shared memory between stages
- Each stage closes its output channel when done (signals downstream)
- Must pass `go test -race`
- No third-party libraries
- No global variables

---

## Function Signatures (suggested, not mandatory)

```go
type Record struct {
    ID       string  `json:"id"`
    Category string  `json:"category"`
    Amount   float64 `json:"amount"`
    Quantity int     `json:"quantity"`
}

type EnrichedRecord struct {
    Record
    Score float64
}

type CategorySummary struct {
    Category   string
    Count      int
    TotalScore float64
    AvgScore   float64
}

func LoadRecords(path string) ([]Record, error)

func StageValidate(ctx context.Context, in <-chan Record, workers int) <-chan Record
func StageEnrich(ctx context.Context, in <-chan Record, workers int) <-chan EnrichedRecord
func StageAggregate(ctx context.Context, in <-chan EnrichedRecord, workers int) <-chan CategorySummary

func PrintSummary(summaries []CategorySummary, dropped int, wallTime time.Duration)
```

---

## Input

`input/records.json` — array of record objects:

```json
[
  { "id": "r1", "category": "electronics", "amount": 299.99, "quantity": 2 },
  ...
]
```

| Field      | Type    | Description                |
|------------|---------|----------------------------|
| `id`       | string  | Unique record ID           |
| `category` | string  | Product category           |
| `amount`   | float64 | Transaction amount         |
| `quantity` | int     | Item count                 |

---

## Expected Output

See `output/expected_output.txt`.

Sample stdout:

```
[validate] PASS r1  electronics  amount=299.99  qty=2
[validate] DROP r5  ""           reason=invalid_category
[enrich]   r1   score=121.60
[aggregate] electronics  count=3  total=364.80  avg=121.60
...

=== Pipeline Summary ===
Records loaded   : 15
Records dropped  : 3
Records processed: 12
Wall time        : ~310ms

Category Breakdown:
  electronics  count=4  avg_score=134.20
  clothing     count=3  avg_score=88.50
  food         count=3  avg_score=45.10
  books        count=2  avg_score=31.80
```

---

## Evaluation Criteria

| Area              | What to look for                                               |
|-------------------|----------------------------------------------------------------|
| Correctness       | Score formula applied right; invalid records dropped           |
| Concurrency       | Stages truly parallel; channels wired correctly                |
| Backpressure      | Slow stage blocks fast stage via buffered channel              |
| Cancellation      | All goroutines exit on ctx cancel; no leaks                    |
| Channel discipline| Each stage owns close of its output channel                    |
| Code structure    | Clean separation — loader, each stage, reporter                |
| Race detector     | `go test -race` passes                                         |

---

## Bonus (optional)

- [ ] CLI flags for worker counts per stage and timeout (`flag` package)
- [ ] Add Stage 4 — **Emit**: write aggregated results to `output/results.json`
- [ ] Unit test for score formula
- [ ] Unit test for validation rules — table-driven
- [ ] Metrics: track per-stage latency (avg time a record spends in each stage)

---

## Setup

```bash
cd go/src/Go_Exercise/multi-stage-pipeline
go mod init multi-stage-pipeline
go run main.go
go test -race ./...
```

---

## Key Go Concepts Exercised

- Pipeline pattern (channel chaining)
- Fan-out worker pools per stage
- `context.WithTimeout` + propagation across stages
- Channel ownership (who closes)
- `sync.WaitGroup` for draining workers before closing output channel
- `sync.Mutex` for aggregation state
- Buffered channels for backpressure
- JSON decoding (`encoding/json`)
- Race detector (`-race`)
