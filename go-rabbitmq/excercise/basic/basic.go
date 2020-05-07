package main

import (
	"log"

	"github.com/streadway/amqp"
	"github.com/dindasigma/go-rabbitmq/excercise/utils"
)

func main() {
	go client()
	go server()

	// to keep main go routine alive while the others are busily publishing and receiving messages
	forever := make(chan bool)
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func client() {
	conn, ch := utils.GetChannel()
	defer conn.Close()
	defer ch.Close()

	q := utils.GetQueue("hello", false, false, false, false, nil, ch)

	msgs, err := ch.Consume (
		q.Name, // name of queue
		"", // consumer
		true, // autoAck
		false, // exclusive
		false, // nolocal
		false, // nowait
		nil, // header
	)

	utils.FailOnError(err, "Failed to register a consumer")

	for msg := range msgs {
		log.Printf("Received message with message: %s", msg.Body)
	}
}

func server() {
	conn, ch := utils.GetChannel()
	defer conn.Close()
	defer ch.Close()

	q := utils.GetQueue("hello", false, false, false, false, nil, ch)

	msg := amqp.Publishing{
		ContentType: "text/plain",
		Body: []byte("Hello RabbitMQ"),
	}

	// publish with direct exchange type (default): guarantee that only one of client get the message and only one time
	ch.Publish("", q.Name, false, false, msg)
}