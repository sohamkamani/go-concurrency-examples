package main

import (
	"fmt"
	"sync"
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

type salad struct {
	sync.Mutex
	vegetables []*vegetable
}

func (s *salad) add(v *vegetable) {
	s.Lock()
	s.vegetables = append(s.vegetables, v)
	s.Unlock()
}

func (s *salad) chopAndAdd(in <-chan *vegetable, out chan<- struct{}) {
	for {
		v := <-in
		v.chop()
		s.add(v)
		out <- struct{}{}
	}
}

func main() {
	tomato := &vegetable{name: "tomato"}
	cucumber := &vegetable{name: "cucumber"}

	in := make(chan *vegetable)
	out := make(chan struct{})

	s := &salad{}
	for i := 0; i < 2; i++ {
		go s.chopAndAdd(in, out)
	}

	in <- tomato
	in <- cucumber

	<-out
	<-out
	fmt.Println("Finished preparing salad with", len(s.vegetables), "vegetables")

}
