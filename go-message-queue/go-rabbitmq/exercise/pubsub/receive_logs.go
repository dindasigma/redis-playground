package main

import (
	"log"

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

	// create a queue with a random name (let the server choose a random queue name for us)
	// and when the connection that declared it closes, the queue will be deleted because it is declared as exclusive.
	q := utils.GetQueue("", false, false, true, false,nil, ch)

	// bind queue and exchange, tell the exchange to send messages to our queue
	err = ch.QueueBind(
		q.Name, // queue name
		"", // routing key
		"logs", // exchange
		false,
		nil,
	)
	utils.FailOnError(err, "Failed to bind a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"", // consumer
		true, // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil, // args
	)
	utils.FailOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf(" [x] %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever
}
