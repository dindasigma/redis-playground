package main

// channel: a way for safe communicate between go routines, so they no need to care about each other, just know the channel
// unbuffered channel: channel without internal some capacity
// step: create channel, send message into it, receive message from it

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	// to create channel, using built-in make function, also have to declare type of message
	ch := make(chan int)

	wg.Add(2)

	go func(ch chan int, wg *sync.WaitGroup) {
		// receive message
		fmt.Println(<-ch)
		wg.Done()
	}(ch, wg)

	go func(ch chan int, wg *sync.WaitGroup) {
		// send message
		ch <- 42
		wg.Done()
	}(ch, wg)

	wg.Wait()
}