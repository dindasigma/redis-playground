package main

// channel: a way for safe communicate between go routines, so they no need to care about each other, just know the channel
// buffered channel: channel with internal some capacity
// step: create channel, send message into it, receive message from it

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	// to create channel, using built-in make function, also have to declare type of message
	ch := make(chan int, 1)

	wg.Add(2)

	go func(ch chan int, wg *sync.WaitGroup) {
		// receive message
		fmt.Println(<-ch)

		// concurrent task has been completed
		wg.Done()
	}(ch, wg)

	go func(ch chan int, wg *sync.WaitGroup) {
		// send message
		ch <- 42
		ch <- 27

		// concurrent task has been completed
		wg.Done()
	}(ch, wg)

	wg.Wait()
}

// result: 42
// we will never get 27 because the internal capacity of message = 1 so it's blocking
