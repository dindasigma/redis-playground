package main

import (
	"fmt"
	"sync"
)

// waits for a collection of goroutines to finish
// add(), wait(), done()

func main() {
	wg := &sync.WaitGroup{}
	id := 1

	// number of concurrent tasks that want to be waiting on
	wg.Add(2)

	go func(id int, wg *sync.WaitGroup) {
		fmt.Println("concurrent task 1")
		// concurrent task has been completed
		wg.Done()
	}(id, wg)

	go func(id int, wg *sync.WaitGroup) {
		fmt.Println("concurrent task 2")
		// concurrent task has been completed
		wg.Done()
	}(id, wg)

	// wait until no longer any concurrent activities
	wg.Wait()
}