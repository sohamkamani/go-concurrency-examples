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

func chopper(in <-chan *vegetable, out chan<- *vegetable) {
	for {
		v := <-in
		v.chop()
		out <- v
	}
}

func main() {
	tomato := &vegetable{name: "tomato"}
	cucumber := &vegetable{name: "cucumber"}

	in := make(chan *vegetable, 2)
	out := make(chan *vegetable, 2)

	for i := 0; i < 2; i++ {
		go chopper(in, out)
	}

	in <- tomato
	in <- cucumber

	fmt.Println("Chopped vegetables:")
	v := <-out
	fmt.Println(v.name, ":", v.chopped)
	v = <-out
	fmt.Println(v.name, ":", v.chopped)

}
