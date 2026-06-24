package main

import "fmt"

type Company struct {
	Name  string
	Price int
}
type Stock struct {
	name   string
	prices []int
}

func main() {
	stocks := []Stock{
		{"AAPL", []int{150, 152, 149, 155, 160}},
		{"GOOG", []int{2800, 2750, 2820, 2795, 2810}},
		{"TSLA", []int{700, 710, 695, 720, 715}},
	}

	N := 10
	done := make(chan bool)

	ch := make(chan Company)
	for _, stock := range stocks {
		go sendSignals(stock, ch, done)
	}

	go func() {
		for i := 0; i < N; i++ {
			s := <-ch
			fmt.Println(s.Name, ":", s.Price)
		}
		close(done)
	}()

	<-done
	fmt.Println("received", N, "updates, stopping")
}

func sendSignals(stock Stock, ch chan Company, done <-chan bool) {
	for _, p := range stock.prices {
		select {
		case ch <- Company{Name: stock.name, Price: p}:
		case <-done:
			return
		}
	}
}
