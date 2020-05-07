package main

// just need to tell which topic to consume

import (
	"log"
	"sync"

	"github.com/nsqio/go-nsq"
)

func main(){
	wg := &sync.WaitGroup{}
	wg.Add(1)

	config := nsq.NewConfig()
	consumer, err := nsq.NewConsumer("test-nsq", "ch", config)
	if err != nil {
		log.Fatal(err)
	}

	// set the Handler for messages received by this Consumer.
	consumer.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		log.Println(string(message.Body))
		return nil
	}))

	err = consumer.ConnectToNSQD("localhost:4150")
	if err != nil {
		log.Fatal(err)
	}

	wg.Wait()
}