package main

import (
	"fmt"
)

// getType ...
func getType(i interface{}) {
	v, ok := i.(int)
	fmt.Println(v, ok)
}

// func main ...
func main() {
	var s interface{} = 100
	getType(s)
}
