package main

import "fmt"

func main() {
	ch := make(chan struct{})
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			go even(i, ch)
		} else {
			go odd(i, ch)
		}
		<-ch
	}
}

func even(i int, ch chan struct{}) {
	fmt.Println("Print ", i)
	ch <- struct{}{}
}
func odd(i int, ch chan struct{}) {
	fmt.Println("Print ", i)
	ch <- struct{}{}
}
