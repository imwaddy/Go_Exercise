package main

import (
	"fmt"
	"runtime"
)

// getGOMAXPROCS ...
func getGOMAXPROCS() int {
	return runtime.GOMAXPROCS(-1)
}

// main func
func main() {
	fmt.Println(getGOMAXPROCS())
}
