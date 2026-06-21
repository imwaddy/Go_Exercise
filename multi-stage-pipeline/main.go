package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
	"sync"
	"time"
)

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

type aggregateState struct {
	mu   sync.Mutex
	data map[string]*CategorySummary
}

func LoadRecords(path string) ([]Record, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var records []Record
	if err := json.Unmarshal(data, &records); err != nil {
		return nil, err
	}
	return records, nil
}

func isInvalid(rec Record) (bool, string) {
	if strings.TrimSpace(rec.ID) == "" {
		return true, "missing_id"
	}
	if rec.Amount <= 0 {
		return true, "amount_must_be_positive"
	}
	if rec.Quantity < 1 {
		return true, "invalid_quantity"
	}
	if !slices.Contains([]string{"electronics", "clothing", "food", "books"}, rec.Category) {
		return true, "invalid_category"
	}
	return false, ""
}

func StageValidate(ctx context.Context, in <-chan Record, workers int) (<-chan Record, <-chan int) {
	out := make(chan Record, 10)
	dropped := make(chan int, 1)

	go func() {
		fmt.Println("Validation log:")
		defer close(out)
		defer close(dropped)
		count := 0

		for rec := range in {
			if invalid, reason := isInvalid(rec); invalid {
				fmt.Printf("[validate] DROP %-4s %-12s reason=%s\n", rec.ID, rec.Category, reason)
				count++
				continue
			}
			fmt.Printf("[validate] PASS %-4s %-12s amount=%.2f qty=%d\n", rec.ID, rec.Category, rec.Amount, rec.Quantity)
			out <- rec
		}
		dropped <- count
	}()
	return out, dropped
}

func StageEnrich(ctx context.Context, in <-chan Record, workers int) <-chan EnrichedRecord {
	out := make(chan EnrichedRecord, 10)

	go func() {
		fmt.Println("\nEnrich log (score = amount*0.4 + quantity*0.6):")
		defer close(out)
		for rec := range in {
			score := (rec.Amount * 0.4) + (float64(rec.Quantity) * 0.6)
			fmt.Printf("[enrich] %-4s score=%.2f\n", rec.ID, score)
			out <- EnrichedRecord{Record: rec, Score: score}
		}
	}()
	return out
}

func StageAggregate(ctx context.Context, in <-chan EnrichedRecord, workers int) <-chan CategorySummary {
	out := make(chan CategorySummary)
	data := make(map[string]*CategorySummary)
	var wg sync.WaitGroup
	var mutex sync.Mutex

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for rec := range in {
				mutex.Lock()
				val, ok := data[rec.Category]
				if !ok {
					data[rec.Category] = &CategorySummary{Category: rec.Category, AvgScore: rec.Score, Count: 1, TotalScore: rec.Score}
				} else {
					val.Count += 1
					val.TotalScore += rec.Score
					val.AvgScore = val.TotalScore / float64(val.Count)
				}
				mutex.Unlock()
			}
		}()
	}

	go func() {
		wg.Wait()
		defer close(out)
		for _, s := range data {
			out <- *s
		}
	}()

	return out
}

func PrintSummary(summaries []CategorySummary, loaded, dropped int, wallTime time.Duration) {
	fmt.Println("\n=== Pipeline Summary ===")
	fmt.Printf("  Records loaded   : %d\n", loaded)
	fmt.Printf("  Records dropped  : %d\n", dropped)
	fmt.Printf("  Records processed: %d\n", loaded-dropped)
	fmt.Printf("  Wall time        : ~%dms\n", wallTime.Milliseconds())
	fmt.Println("\nCategory Breakdown:")
	for _, s := range summaries {
		fmt.Printf("  %-12s count=%-3d total_score=%-8.2f avg_score=%.2f\n",
			s.Category, s.Count, s.TotalScore, s.AvgScore)
	}
}

func main() {
	records, err := LoadRecords("input/records.json")
	if err != nil {
		log.Fatalf("load: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	in := make(chan Record, len(records))
	for _, r := range records {
		in <- r
	}
	close(in)

	start := time.Now()

	validated, dropped := StageValidate(ctx, in, 2)
	enriched := StageEnrich(ctx, validated, 3)
	aggregated := StageAggregate(ctx, enriched, 2)

	var summaries []CategorySummary
	for s := range aggregated {
		summaries = append(summaries, s)
	}

	totalDropped := 0
	for d := range dropped {
		totalDropped += d
	}

	PrintSummary(summaries, len(records), totalDropped, time.Since(start))
}
