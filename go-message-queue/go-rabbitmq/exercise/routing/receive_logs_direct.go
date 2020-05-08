package main

import (
	"log"
	"os"

	"github.com/dindasigma/go-rabbitmq/exercise/utils"
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

	// create a queue with a random name (let the server choose a random queue name for us)
	// and when the connection that declared it closes, the queue will be deleted because it is declared as exclusive.
	q := utils.GetQueue("", false, false, true, false,nil, ch)

	if len(os.Args) < 2 {
		log.Printf("Usage: %s [info] [warning] [error]", os.Args[0])
		os.Exit(0)
	}

	// bind exchange with multiple queues
	for _, s := range os.Args[1:] {
		log.Printf("Binding queue %s to exchange %s with routing key %s",
			q.Name, "logs_direct", s)
		err = ch.QueueBind(
			q.Name, // queue name
			s, // routing key
			"logs_direct", // exchange
			false,
			nil)
		utils.FailOnError(err, "Failed to bind a queue")
	}

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