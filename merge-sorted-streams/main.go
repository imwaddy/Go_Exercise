package main

import (
	"container/heap"
	"context"
	"fmt"
	"time"
)

type Item struct {
	value int
	ch    <-chan int
}

type MinHeap []Item

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].value < h[j].value }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x any)        { *h = append(*h, x.(Item)) }
func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func main() {
	streams := [][]int{
		{1, 4, 7, 10},
		{2, 5, 8},
		{3, 6, 9, 12},
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(30*time.Second))
	defer cancel()

	var chs []<-chan int
	for _, s := range streams {
		chs = append(chs, streamGenerator(ctx, s))
	}

	for v := range merge(ctx, chs...) {
		fmt.Println(v)
	}
}

func streamGenerator(ctx context.Context, nums []int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for _, n := range nums {
			select {
			case ch <- n:
			case <-ctx.Done():
				return
			}
		}
	}()
	return ch
}

func merge(ctx context.Context, streams ...<-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		h := &MinHeap{}
		heap.Init(h)

		// seed heap — first value from each stream
		for _, ch := range streams {
			select {
			case v, ok := <-ch:
				if ok {
					heap.Push(h, Item{value: v, ch: ch})
				}
			case <-ctx.Done():
				return
			}
		}

		for h.Len() > 0 {
			item := heap.Pop(h).(Item)

			select {
			case out <- item.value:
			case <-ctx.Done():
				return
			}

			// refill from same channel
			select {
			case v, ok := <-item.ch:
				if ok {
					heap.Push(h, Item{value: v, ch: item.ch})
				}
			case <-ctx.Done():
				return
			}
		}
	}()

	return out
}
