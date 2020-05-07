package main

// tasks will survive even if RabbitMQ is restarted.

// add d.Ack(false) (this acknowledges a single delivery), once we're done with a task.
// set durable = true on QueueDeclare on both producer and consumer
// mark our messages as persistent - by using the amqp.Persistent option amqp.Publishing takes
// set the prefetch count on worker with the value of 1. This tells RabbitMQ not to give more than one message to a worker at a time. Or, in other words, don't dispatch a new message to a worker until it has processed and acknowledged the previous one. Instead, it will dispatch it to the next worker that is not still busy.

import (
	"github.com/dindasigma/go-rabbitmq/excercise/utils"
	"github.com/streadway/amqp"
	"log"
	"os"
	"strings"
)

func main() {
	conn, ch := utils.GetChannel()
	defer conn.Close()
	defer ch.Close()

	q := utils.GetQueue("task_queue", true, false, false, false,nil, ch)

	body := bodyFrom(os.Args)
	msg := amqp.Publishing{
		DeliveryMode: amqp.Persistent, //  queue won't be lost even if RabbitMQ restarts
		ContentType: "text/plain",
		Body: []byte(body),
	}
	err := ch.Publish(
		"", // exchange
		q.Name, // routing key
		false, // mandatory
		false,
		msg)
	utils.FailOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s", body)
}

func bodyFrom(args []string) string {
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		s = "hello"
	} else {
		s = strings.Join(args[1:], " ")
	}

	return s
}
