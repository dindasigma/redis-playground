package main

// avoid doing a resource-intensive task immediately and having to wait for it to complete

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
