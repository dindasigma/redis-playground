package main

// another client
// receive message from queue "hello" that publishing on basic.go
// just to ensure the direct exchange type

import (
	"log"

	"github.com/dindasigma/go-rabbitmq/excercise/utils"
)

func main() {
	conn, ch := utils.GetChannel()
	defer conn.Close()
	defer ch.Close()

	q := utils.GetQueue("hello", ch)

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	utils.FailOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}