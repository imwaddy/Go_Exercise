# Container With Most Water

Given array `heights` where `heights[i]` = height of vertical line at index `i`.
Find two lines forming container that holds maximum water. Return the area.

```
heights = [1, 8, 6, 2, 5, 4, 8, 3, 7]

  8         8
  |     5   |   7
  | 6   |   | 3 |
  | | 2 | 4 | | |
1 | | | | | | | |
─────────────────────
0 1 2 3 4 5 6 7 8

Answer: 49  (lines at index 1 and 8, min(8,7)*7 = 49)
```

## Signature

```go
func maxWater(heights []int) int
```

## Constraints

- Two pointer approach — O(n) time, O(1) space
- No brute force O(n²)

## Hint

Start `left=0`, `right=len-1`. Move the pointer with **smaller** height inward.
Think: why move the smaller one, not the larger?

## Files

- `input/input.json` — 4 test cases
- `output/expected_output.txt` — expected answers

## Run

```bash
go run main.go
```
