package main

import (
	"fmt"
	"strconv"
)

// LivingCreature ...
type LivingCreature interface {
	Walk() string
}

// Man ...
type Man struct {
	Legs int
}

// Dog ...
type Dog struct {
	Legs int
}

// Walk ...
func (m *Man) Walk() string {
	return "Man walked on " + strconv.Itoa(m.Legs) + " legs"
}

// Walk ..
func (d *Dog) Walk() string {
	return "Man walked on " + strconv.Itoa(d.Legs) + " legs"
}

// func main
func main() {
	man := Man{
		Legs: 2,
	}
	fmt.Println("Man ==", man.Walk())
	fmt.Println("---------------------------------")
	dog := Dog{
		Legs: 4,
	}
	fmt.Println("Dog ==", dog.Walk())
}
