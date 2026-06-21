# Problem: Concurrent Priority Task Scheduler

## Problem Statement

Build a concurrent task scheduler in Go that reads tasks from a JSON file and executes them using a fixed-size worker pool, respecting task priority.

---

## Requirements

### Core

1. **Read** tasks from `input/tasks.json`.
2. **Priority queue** — tasks with lower `priority` value execute first (1 = highest). Same priority → FIFO.
3. **Worker pool** — fixed number of workers (configurable, default: 3). Workers pick the highest-priority available task.
4. **Simulate work** — each task "runs" for `duration_ms` milliseconds.
5. **Context cancellation** — accept an overall timeout (default: 3s). If timeout fires, cancel all in-flight tasks and drain the queue gracefully.
6. **Result collection** — after all workers finish, print a summary:
   - Which tasks completed, which were cancelled (due to timeout).
   - Total wall-clock time.

### Constraints

- Must use goroutines + channels (no `sync.WaitGroup` as the *only* mechanism — you may use it alongside channels).
- Priority queue must be implemented using `container/heap` (no third-party libs).
- No global variables.
- Handle the case where `duration_ms` of a single task exceeds the overall timeout.

---

## Function Signatures (suggested, not mandatory)

```go
type Task struct {
    ID         string `json:"id"`
    Name       string `json:"name"`
    Priority   int    `json:"priority"`
    DurationMs int    `json:"duration_ms"`
}

type Result struct {
    Task      Task
    WorkerID  int
    Status    string // "completed" | "cancelled"
    StartedAt time.Time
    EndedAt   time.Time
}

// LoadTasks reads and parses input/tasks.json
func LoadTasks(path string) ([]Task, error)

// NewScheduler returns a configured scheduler
func NewScheduler(workers int, timeout time.Duration) *Scheduler

// Run executes all tasks and returns results
func (s *Scheduler) Run(tasks []Task) []Result

// PrintSummary writes the execution report to stdout
func PrintSummary(results []Result, wallTime time.Duration)
```

---

## Input

`input/tasks.json` — array of task objects:

```json
[
  { "id": "t1", "name": "Send Email", "priority": 3, "duration_ms": 200 },
  ...
]
```

| Field         | Type   | Description                          |
|---------------|--------|--------------------------------------|
| `id`          | string | Unique task ID                       |
| `name`        | string | Human-readable name                  |
| `priority`    | int    | 1 (highest) → 5 (lowest)             |
| `duration_ms` | int    | Simulated work duration (ms)         |

---

## Expected Output

See `output/expected_output.txt` for reference execution trace.

Sample stdout:

```
[worker-1] STARTED  t2  "Generate Report"   priority=1
[worker-2] STARTED  t6  "Backup Files"      priority=1
[worker-3] STARTED  t3  "Sync Database"     priority=2
[worker-3] DONE     t3  "Sync Database"     300ms
[worker-1] DONE     t2  "Generate Report"   500ms
...

=== Summary ===
Total tasks    : 10
Completed      : 10
Cancelled      : 0
Wall time      : 912ms
```

---

## Evaluation Criteria

| Area                  | What to look for                                              |
|-----------------------|---------------------------------------------------------------|
| Correctness           | Tasks run in priority order; FIFO tiebreak                    |
| Concurrency           | Workers truly parallel; no race conditions (`go test -race`)  |
| Cancellation          | Context timeout propagates; no goroutine leak                 |
| Code structure        | Clean separation — loader, scheduler, worker, reporter        |
| heap usage            | `container/heap` interface implemented correctly              |
| Error handling        | JSON parse errors, empty input, zero workers handled          |

---

## Bonus (optional)

- [ ] Add retry logic: tasks that fail (simulate random failure) retry up to 2 times before marking failed.
- [ ] Expose worker count and timeout as CLI flags (`flag` package).
- [ ] Write a unit test for the priority queue: verify pop order matches expected priority sequence.
- [ ] Add a rate limiter: max N tasks started per second across all workers.

---

## Setup

```bash
cd go/src/Go_Exercise/priority-task-scheduler
go mod init priority-task-scheduler   # if needed
go run main.go
go test -race ./...
```

---

## Key Go Concepts Exercised

- `container/heap` (priority queue)
- Goroutines + channels (fan-out worker pool)
- `context.WithTimeout` + cancellation propagation
- `sync.WaitGroup` + `sync.Mutex`
- JSON decoding (`encoding/json`)
- `time.Sleep` as work simulation
- Race detector (`-race`)
