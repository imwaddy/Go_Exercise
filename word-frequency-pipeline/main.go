package main

import (
	"fmt"
	"sort"
	"strings"
	"sync"
)

func main() {
	sentences := []string{
		"the cat sat on the mat",
		"the cat sat on the hat",
		"the cat in the hat",
		"the fat cat sat",
		"cat cat cat hat",
	}

	sentence := distribute(sentences)
	words := split(sentence)

	res := make(chan map[string]int)
	n := 0
	for w := range words {
		go count(w, res)
		n++
	}

	freq := make(map[string]int)
	for i := 0; i < n; i++ {
		r := <-res
		for word, c := range r {
			freq[word] += c
		}
	}

	printSummary(freq)

}

func count(words []string, res chan map[string]int) {
	ch := make(map[string]int)

	var m sync.Mutex
	go func() {
		for _, word := range words {
			m.Lock()
			ch[word] += 1
			m.Unlock()
		}
		res <- ch
	}()
}

func distribute(sentences []string) <-chan string {
	ch := make(chan string)
	go func() {
		defer close(ch)
		for _, sentence := range sentences {
			ch <- sentence
		}
	}()
	return ch
}

func split(sentence <-chan string) <-chan []string {
	ch := make(chan []string)
	go func() {
		defer close(ch)
		for s := range sentence {
			words := strings.Split(s, " ")
			ch <- words
		}
	}()
	return ch
}

type WordFrequency struct {
	word  string
	count int
}

func printSummary(freq map[string]int) {
	var pairs []WordFrequency
	for k, v := range freq {
		pairs = append(pairs, WordFrequency{word: k, count: v})
	}

	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i] != pairs[j] {
			return pairs[i].count > pairs[j].count
		}
		return pairs[i].word < pairs[j].word
	})

	for _, p := range pairs[:5] {
		fmt.Println(p.word, ">>", p.count)
	}
}
