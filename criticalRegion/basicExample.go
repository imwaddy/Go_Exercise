package main

import (
	"fmt"
	"time"
)

var global int64

func main() {

	for i := 0; i < 5; i++ {
		go increment()
		go display()
	}
	time.Sleep(time.Second * 5)
}

func increment() {
	global++
}

func display() {
	fmt.Println("Value is ", global)
}
