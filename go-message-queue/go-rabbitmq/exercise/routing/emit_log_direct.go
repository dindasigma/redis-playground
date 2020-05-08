package main

// subscribe only to a subset of the messages
// Bindings can take an extra routing_key parameter
// It is perfectly legal to bind multiple queues with the same binding key
// will be using direct exchange

import (
	"log"
	"os"

	"github.com/dindasigma/go-rabbitmq/exercise/utils"
	"github.com/streadway/amqp"
)


func main() {
	conn, ch := utils.GetChannel()
	defer conn.Close()
	defer ch.Close()

	err := ch.ExchangeDeclare(
		"logs_direct", // name
		"direct", // type
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
		"logs_direct", // exchange
		utils.SeverityFrom(os.Args), // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	utils.FailOnError(err, "Failed to publish a message")

	log.Printf(" [x] Sent %s", body)
}