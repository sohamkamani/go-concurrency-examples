package main

import (
	"fmt"
	"sync"
	"time"
)

type vegetable struct {
	name    string
	chopped bool
	// Add a waitgroup pointer attribute
	wg *sync.WaitGroup
}

func (v *vegetable) chop() {
	// We use the `Add` method to add 1 member to the waitgroup
	v.wg.Add(1)
	time.Sleep(50 * time.Millisecond)
	v.chopped = true
	// Once we're done, we call the done method to release the
	// member of the waitgroup
	v.wg.Done()
}

func main() {
	// initialize our vegetables
	var wg sync.WaitGroup
	tomato := &vegetable{name: "tomato", wg: &wg}
	cucumber := &vegetable{name: "cucumber", wg: &wg}

	go tomato.chop()
	go cucumber.chop()

	// The `Wait` method waits for all members of the waitgroup
	// to be released
	wg.Wait()

	fmt.Println("Chopped vegetables:")
	fmt.Println("tomato:", tomato.chopped)
	fmt.Println("cucumber:", cucumber.chopped)
}
