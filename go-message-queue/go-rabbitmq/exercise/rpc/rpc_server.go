package main

// go run rpc_server.go
// go run rpc_client.go 30

import (
	"log"
	"strconv"

	"github.com/streadway/amqp"
	"github.com/dindasigma/go-rabbitmq/exercise/utils"
)

func main() {
	conn, ch := utils.GetChannel()
	defer conn.Close()
	defer ch.Close()

	q := utils.GetQueue("rpc_queue", false, false, true, false,nil, ch)

	err := ch.Qos(
		1, // prefetch count
		0, // prefetch size
		false, // global
	)
	utils.FailOnError(err, "Failed to set QoS")

	msgs, err := ch.Consume(
		q.Name, // queue
		"", // consumer
		false, // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil, // args
	)
	utils.FailOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			n, err := strconv.Atoi(string(d.Body))
			utils.FailOnError(err, "Failed to convert body to integer")

			log.Printf(" [.] fib(%d)", n)
			response := fib(n)

			err = ch.Publish(
				"", // exchange
				d.ReplyTo, // routing key
				false, // mandatory
				false, // immediate
				amqp.Publishing{
					ContentType: "text/plain",
					CorrelationId: d.CorrelationId,
					Body: []byte(strconv.Itoa(response)),
				})
			utils.FailOnError(err, "Failed to publish a message")

			d.Ack(false)
		}
	}()

	log.Printf(" [*] Awaiting RPC requests")
	<-forever
}

func fib(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fib(n-1) + fib(n-2)
	}
}