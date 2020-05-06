package main

import (
	"fmt"
	"sync"
)

// there may came a time when we created a channel, then we no longer have messages to send
// we cannot check for closed channel,
// when we sending new messages into it, will triggers a panic. but receiving messages is okay

func main() {
	wg := &sync.WaitGroup{}
	ch := make(chan int, 1)

	wg.Add(2)
	go func(ch chan int, wg *sync.WaitGroup) {

		// check channel, will return ok=true if channel open
		if msg, ok := <-ch; ok {
			fmt.Println(msg, ok) // return 42, true
		}
		close(ch)
		fmt.Println(<-ch) // return 0 value but no error

		wg.Done()
	}(ch, wg)

	go func(ch chan int, wg *sync.WaitGroup) {
		ch <- 42

		// close(ch)
		// ch <- 42 // panic: send on closed channel

		wg.Done()
	}(ch, wg)

	wg.Wait()
}