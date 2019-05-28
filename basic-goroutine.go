package main

import (
	"fmt"
	"time"
)

type vegetable struct {
	name    string
	chopped bool
}

func (v *vegetable) chop() {
	time.Sleep(50 * time.Millisecond)
	v.chopped = true
}

func main() {
	// initialize our vegetables
	tomato := &vegetable{name: "tomato"}
	cucumber := &vegetable{name: "cucumber"}

	// Call tomato's chop method as a goroutine
	go tomato.chop()
	// Cucumbers chop method blocks
	cucumber.chop()

	fmt.Println("Chopped vegetables:")
	fmt.Println("tomato:", tomato.chopped)
	fmt.Println("cucumber:", cucumber.chopped)
}
