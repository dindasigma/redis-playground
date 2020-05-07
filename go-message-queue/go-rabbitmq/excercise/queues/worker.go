package main

import (
	"bytes"
	"log"
	"time"

	"github.com/dindasigma/go-rabbitmq/excercise/utils"
)

func main() {
	conn, ch := utils.GetChannel()
	defer conn.Close()
	defer ch.Close()

	// queue durable: true ->  queue won't be lost even if RabbitMQ restarts
	q := utils.GetQueue("task_queue", true, false, false, false, nil, ch)

	//
	err := ch.Qos(
		1, // prefetch count: not to give more than one message to a worker at a time
		0,     // prefetch size
		false, // global
	)
	utils.FailOnError(err, "Failed to set QoS")

	msgs, err := ch.Consume(
		q.Name, // queue
		"", // consumer
		false, // auto-ack, we will use manual message acknowledgements
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil, // args
	)
	utils.FailOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)

			// every dot will account for one second of "work". For example, a fake task described by Hello... will take three seconds
			dot_count := bytes.Count(d.Body, []byte("."))
			t := time.Duration(dot_count)
			time.Sleep(t * time.Second)
			log.Print("Done")
			// send a proper acknowledgment with multiple:false from the worker once we're done with a task.
			d.Ack(false)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}