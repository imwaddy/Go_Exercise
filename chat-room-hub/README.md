# Exercise: Concurrent Chat Room Hub

**Difficulty:** Beginner → Intermediate  
**Concepts:** goroutines, channels, fan-in, fan-out, `select`, done channel, WaitGroup

---

## Problem Statement

Build a simple chat room hub where multiple users send messages concurrently and the hub broadcasts each message to all other users.

```
User A ──┐
User B ──┤──► Hub (broadcast) ──► all users receive
User C ──┘
```

---

## Part 1 — Beginner: Fan-In Collector

**Goal:** Merge messages from N user goroutines into one channel.

Each user goroutine sends exactly `M` messages in the format `"<username>: <message>"`.  
Main reads from the merged channel and prints every message.

### Requirements

1. Each user runs in its own goroutine
2. Single output channel receives all messages (fan-in)
3. No `time.Sleep` for ordering
4. Output channel closes after **all** users finish — no deadlock
5. Main must not exit before all messages are printed

### Function Signatures

```go
func userSender(name string, messages []string, out chan<- string, wg *sync.WaitGroup)
func fanIn(users []<-chan string) <-chan string
```

### Sample Input

```go
users := map[string][]string{
    "alice": {"hello", "how are you"},
    "bob":   {"hey", "doing well"},
}
```

### Sample Output (order may vary)

```
alice: hello
bob: hey
alice: how are you
bob: doing well
```

### Things to figure out

- Who closes the output channel and when?
- Why does fan-in need a separate goroutine per input channel?
- What happens without `WaitGroup`?

---

## Part 2 — Intermediate: Hub with Broadcast

**Goal:** Hub receives messages and fan-outs to all registered subscribers.

### Requirements

1. `Hub` struct with:
   - `subscribe(name string) <-chan string` — returns channel for that user to receive messages
   - `publish(msg string)` — sends message to all subscribers except sender
   - `shutdown()` — closes all subscriber channels cleanly
2. N goroutines subscribe, then M goroutines publish concurrently
3. No subscriber misses any message published after they subscribed
4. `shutdown()` must not deadlock even if a subscriber stops reading

### Struct Skeleton

```go
type Hub struct {
    // your fields
}

func NewHub() *Hub
func (h *Hub) Subscribe(name string) <-chan string
func (h *Hub) Publish(sender, msg string)
func (h *Hub) Shutdown()
```

### Sample Scenario

```
Alice subscribes → gets channel chA
Bob subscribes   → gets channel chB

Alice publishes "hello"  → Bob receives "alice: hello"
Bob publishes "hi back"  → Alice receives "bob: hi back"

Shutdown → both channels close → both readers exit range loop
```

### Things to figure out

- How do you protect the subscriber list from concurrent writes?
- What buffered channel size prevents a slow reader from blocking the hub?
- How does `Shutdown` signal all readers without knowing who's still reading?

---

## Part 3 — Stretch: Timeout + Select

**Goal:** Add per-message timeout. If hub can't deliver to a subscriber within 100ms, drop the message and log a warning.

### Requirements

1. Use `select` with `time.After` for non-blocking send to each subscriber
2. Dropped messages increment a `DroppedCount` counter (atomic)
3. Add a `Stats()` method returning total published, total dropped

```go
func (h *Hub) Stats() (published, dropped int64)
```

---

## Progression Map

| Part | New concept introduced |
|------|----------------------|
| 1 | goroutines, channels, fan-in, WaitGroup |
| 2 | fan-out, mutex, buffered channels, clean shutdown |
| 3 | `select`, `time.After`, atomic counters |

---

## Verification

```bash
go run main.go
go test -race ./...
```

### Race detector is mandatory — all parts must pass `-race` clean.

---

## Follow-Up

- Add a `History(n int) []string` method returning last N messages
- Support private messages: `@bob hello` routes only to Bob
- Add context cancellation: `NewHubWithContext(ctx context.Context)`
