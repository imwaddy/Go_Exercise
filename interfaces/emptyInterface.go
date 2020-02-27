package main

import (
	"fmt"
)

// CaptureValue ...
func CaptureValue(i interface{}) {
	fmt.Println(i)
}

// func main
func main() {
	CaptureValue("Mayur")
	CaptureValue(50)
	data := struct {
		name string
	}{
		name: "Mayur",
	}
	CaptureValue(data)
}
