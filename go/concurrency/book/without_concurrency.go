package book

import (
	"fmt"
	"math/rand"
	"time"
)

var cacheWithoutConcurrency = map[int]Book{}
var rndWithoutConcurrency = rand.New(rand.NewSource(time.Now().UnixNano()))

func WithoutConcurrency() {
	fmt.Println("without concurrency")

	for i := 0; i < 10; i++ {
		id := rndWithoutConcurrency.Intn(10) + 1

		if b, ok := queryCache(id, cacheWithoutConcurrency); ok {
			fmt.Println("from cache")
			fmt.Println(b)
			continue
		}

		if b, ok := queryDatabase(id, cacheWithoutConcurrency); ok {
			fmt.Println("from database")
			fmt.Println(b)
			continue
		}

		fmt.Printf("Book not found with id: '%v'", id)
		// time.Sleep(150 * time.Millisecond)
	}
}
