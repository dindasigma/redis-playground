package main

import (
	"fmt"
	"math/rand"
	"time"
	"sync"
)

var cache = map[int]Book{}
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func withConcurrency() {
	fmt.Println("with concurrency")

	wg := &sync.WaitGroup{}
	m := &sync.RWMutex{}
	cacheCh := make(chan Book)
	dbCh := make(chan Book)

	for i := 0; i < 10; i++ {
		id := rnd.Intn(10) + 1
		wg.Add(2)

		go func(id int, wg *sync.WaitGroup, m *sync.RWMutex, ch chan<- Book) {
			if b, ok := queryCacheWithConcurrency(id); ok {
				ch <- b
			}
			wg.Done()
		}(id, wg, m, cacheCh)

		go func(id int, wg *sync.WaitGroup, m *sync.RWMutex, ch chan<- Book) {
			if b, ok := queryDatabaseWithConcurrency(id); ok {
				m.Lock()
				cache[id] = b
				m.Unlock()
				ch <- b
				
			}
		}(id, wg, m, dbCh)

		// go routine to handle response
		go func(cacheCh, dbCh <-chan Book) {
			select {
			case b := <-cacheCh:
				fmt.Println("from cache")
				fmt.Println(b)
				<-dbCh
			case b := <-dbCh:
				fmt.Println("from database")
				fmt.Println(b)
			}
		}(cacheCh, dbCh)
		time.Sleep(150 * time.Millisecond)
	}

	wg.Wait()
}

func queryCacheWithConcurrency(id int) (Book, bool) {
	b, ok := cache[id]
	return b, ok
}

func queryDatabaseWithConcurrency(id int) (Book, bool) {
	for _, b := range books {
		if b.ID == id {
			cache[id] = b
			return b, true
		}
	}
	return Book{}, false
}