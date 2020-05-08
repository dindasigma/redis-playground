package main

// just need to tell which topic to consume

import (
	"errors"
	"log"
	"sync"

	"github.com/nsqio/go-nsq"
)

func main(){
	wg := &sync.WaitGroup{}
	wg.Add(1)

	config := nsq.NewConfig()
	consumer, err := nsq.NewConsumer("clicks", "metrics", config)
	if err != nil {
		log.Fatal(err)
	}

	// set the Handler for messages received by this Consumer.
	consumer.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		log.Println(string(message.Body))
		// Returning nil will automatically send a FIN command to NSQ to mark the message as processed.
		return nil

		// Returning a non-nil error will automatically send a REQ command to NSQ to re-queue the message.
		return errors.New("sorry")
	}))

	err = consumer.ConnectToNSQD("localhost:4150")
	if err != nil {
		log.Fatal(err)
	}

	wg.Wait()
}