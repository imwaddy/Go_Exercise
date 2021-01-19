package main

import (
	"fmt"
	"sync"
)

var global int64

var mutex sync.Mutex

var wg sync.WaitGroup

func main() {

	for i := 0; i < 10; i++ {
		wg.Add(2)
		go increment()
		go display()
	}
	wg.Wait()
}

func increment() {
	mutex.Lock()
	global++
	mutex.Unlock()
	wg.Done()
}

func display() {
	mutex.Lock()
	fmt.Println("Value is ", global)
	mutex.Unlock()
	wg.Done()
}
