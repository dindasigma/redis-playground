package main

import (
	"fmt"

	"github.com/dindasigma/redis-playground/packages"
)

var client = playground.Server{}

func main() {
	fmt.Println("Go Redis Playground")
	client.Init()
	client.Basic()
	client.Pubsub()
}