package main

// deliver a message to multiple consumers
// using own declared exchanges: receives messages from producers and the other side it pushes them to queues
// exchange type: fanout (broadcasts all the messages it receives to all the queues it knows)
// by default exchange type is direct (guarantee that only one of client get the message and only one time)

// so we will publish to exchange not queue

import (
	"log"
	"os"

	"github.com/streadway/amqp"
	"github.com/dindasigma/go-rabbitmq/exercise/utils"
)

func main() {
	conn, ch := utils.GetChannel()
	defer conn.Close()
	defer ch.Close()

	err := ch.ExchangeDeclare(
		"logs", // name
		"fanout", // type
		true, // durable
		false, // auto-deleted
		false, // internal
		false, // no-wait
		nil, // arguments
	)
	utils.FailOnError(err, "Failed to declare an exchange")

	body := utils.BodyFrom(os.Args)

	// publish with our logs exchange
	err = ch.Publish(
		"logs", // exchange
		"",     // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	utils.FailOnError(err, "Failed to publish a message")

	log.Printf(" [x] Sent %s", body)
}

