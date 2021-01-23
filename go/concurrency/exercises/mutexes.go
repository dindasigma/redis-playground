package main

import (
	"fmt"
	"sync"
)

// mutex: a mutual exclusion lock
// manage/protect shared memory, so only one task or the owner of Mutexes can access that code

func main() {
	wg := &sync.WaitGroup{}
	m := &sync.Mutex{}

	id := 1
	v := make(map[int]int)

	wg.Add(2)

	go func(id int, wg *sync.WaitGroup, m *sync.Mutex) {
		m.Lock() // Lock so only one goroutine at a time can access map v
		v[id]++
		m.Unlock() // unlock
		wg.Done()
	}(id, wg, m)

	go func(id int, wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Printf("concurrent task 2 with %d", v)
		wg.Done()
	}(id, wg, m)

	// wait until no longer any concurrent activities
	wg.Wait()
}
