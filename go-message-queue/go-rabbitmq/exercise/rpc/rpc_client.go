package main

import (
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/streadway/amqp"
	"github.com/dindasigma/go-rabbitmq/exercise/utils"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	n := utils.BodyFrom(os.Args)
	i, _ := strconv.Atoi(n)

	log.Printf(" [x] Requesting fib(%d)", n)
	res, err := fibonacciRPC(i)
	utils.FailOnError(err, "Failed to handle RPC request")

	log.Printf(" [.] Got %d", res)
}

func randomString(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(randInt(65, 90))
	}
	return string(bytes)
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func fibonacciRPC(n int) (res int, err error) {
	conn, ch := utils.GetChannel()
	defer conn.Close()
	defer ch.Close()

	q := utils.GetQueue("", false, false, true, false,nil, ch)

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

	corrId := randomString(32)

	err = ch.Publish(
		"", // exchange
		"rpc_queue", // routing key
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			CorrelationId: corrId,
			ReplyTo: q.Name,
			Body: []byte(strconv.Itoa(n)),
		})
	utils.FailOnError(err, "Failed to publish a message")

	for d := range msgs {
		if corrId == d.CorrelationId {
			res, err = strconv.Atoi(string(d.Body))
			utils.FailOnError(err, "Failed to convert body to integer")
			break
		}
	}

	return
}