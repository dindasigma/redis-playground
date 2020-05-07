package utils

import (
	"fmt"
	"log"
	"os"
	"strings"

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

func GetQueue(name string, durable bool, autoDelete bool, exclusive bool, noWait bool, args amqp.Table, ch *amqp.Channel) *amqp.Queue {
	q, err := ch.QueueDeclare(name, durable, autoDelete, exclusive, noWait, args)
	FailOnError(err, "Failed to declare queue")

	return &q
}

func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func BodyFrom(args []string) string {
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		s = "hello"
	} else {
		s = strings.Join(args[1:], " ")
	}

	return s
}