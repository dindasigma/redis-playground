package main

// created channels are always bidirectional
// bidirectional: func myFunction(ch chan int) {}
// send-only channel: func myFunction(ch chan<- int) {}
// receive-only channel: func myFunction(ch <-chan int) {}

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	// to create channel, using built-in make function, also have to declare type of message
	ch := make(chan int, 1)

	wg.Add(2)

	go func(ch <-chan int, wg *sync.WaitGroup) {
		// receive message
		fmt.Println(<-ch)
		// ch <- 42 // will error

		// concurrent task has been completed
		wg.Done()
	}(ch, wg)

	go func(ch chan<- int, wg *sync.WaitGroup) {
		// send message
		ch <- 42
		// fmt.Println(<-ch) // will error

		// concurrent task has been completed
		wg.Done()
	}(ch, wg)

	// wait until no longer any concurrent activities
	wg.Wait()
}
