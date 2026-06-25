package main

import (
	"fmt"
	"sync"
)

type Result struct {
	Name    string
	Message string
}

func main() {
	users := map[string][]string{
		"alice": {"hello", "how are you"},
		"bob":   {"hey", "doing well"},
	}
	var wg sync.WaitGroup
	ch := make(chan Result)

	for name, msgs := range users {
		wg.Add(1)
		go userSender(name, msgs, ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for msg := range ch {
		fmt.Println(msg.Name, ":", msg.Message)
	}
}

func userSender(name string, messages []string, out chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, message := range messages {
		out <- Result{Name: name, Message: message}
	}
}

// AI solution
// type Result struct {
// 	Name    string
// 	Message string
// }

// func main() {
// 	users := map[string][]string{
// 		"alice": {"hello", "how are you"},
// 		"bob":   {"hey", "doing well"},
// 	}

// 	var perUser []<-chan Result
// 	for name, msgs := range users {
// 		perUser = append(perUser, userSender(name, msgs))
// 	}

// 	for msg := range fanIn(perUser) {
// 		fmt.Println(msg.Name, ":", msg.Message)
// 	}
// }

// // userSender owns its own channel — creates, fills, closes it
// func userSender(name string, messages []string) <-chan Result {
// 	out := make(chan Result)
// 	go func() {
// 		defer close(out)
// 		for _, m := range messages {
// 			out <- Result{Name: name, Message: m}
// 		}
// 	}()
// 	return out
// }

// // fanIn merges N input channels into one output channel
// func fanIn(channels []<-chan Result) <-chan Result {
// 	var wg sync.WaitGroup
// 	merged := make(chan Result)

// 	forward := func(ch <-chan Result) {
// 		defer wg.Done()
// 		for v := range ch {
// 			merged <- v
// 		}
// 	}

// 	wg.Add(len(channels))
// 	for _, ch := range channels {
// 		go forward(ch)
// 	}

// 	go func() {
// 		wg.Wait()
// 		close(merged)
// 	}()

// 	return merged
// }
