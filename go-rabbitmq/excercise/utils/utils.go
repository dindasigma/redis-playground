package utils

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func GetChannel() (*amqp.Connection, *amqp.Channel) {
	var url = "amqp://guest:guest@localhost"
	conn, err := amqp.Dial(url)
	FailOnError(err, "Failed to establish connection to message broker")

	ch, err := conn.Channel()
	FailOnError(err, "Failed to get channel for connection")

	return conn, ch
}

func GetQueue(name string, ch *amqp.Channel) *amqp.Queue {
	q, err := ch.QueueDeclare(name, false, false, false, false, nil)
	FailOnError(err, "Failed to declare queue")

	return &q
}

func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}
