package main

import (
	"fmt"
	"math/rand"
	"time"
)

var cacheWithoutConcurrency = map[int]Book{}
var rndWithoutConcurrency = rand.New(rand.NewSource(time.Now().UnixNano()))

func withoutConcurrency() {
	fmt.Println("without concurrency")
	
	for i := 0; i < 10; i++ {
		id := rndWithoutConcurrency.Intn(10) + 1
		
		if b, ok := queryCacheWithoutConcurrency(id); ok {
			fmt.Println("from cache")
			fmt.Println(b)
			continue
		}

		if b, ok := queryDatabaseWithoutConcurrency(id); ok {
			fmt.Println("from database")
			fmt.Println(b)
			continue
		}

		fmt.Printf("Book not found with id: '%v'", id)
		// time.Sleep(150 * time.Millisecond)
	}
}

func queryCacheWithoutConcurrency(id int) (Book, bool) {
	b, ok := cacheWithoutConcurrency[id]
	return b, ok
}

func queryDatabaseWithoutConcurrency(id int) (Book, bool) {
	// time.Sleep(100 * time.Millisecond)
	for _, b := range books {
		if b.ID == id {
			cacheWithoutConcurrency[id] = b
			return b, true
		}
	}
	return Book{}, false
}