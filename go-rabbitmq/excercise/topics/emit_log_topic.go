package main

// topic exchange: send messages based on some criteria
// go run receive_logs_topic.go "kern.*" "*.critical"

import (
	"log"
	"os"

	"github.com/streadway/amqp"
	"github.com/dindasigma/go-rabbitmq/excercise/utils"
)

func main() {
	conn, ch := utils.GetChannel()
	defer conn.Close()
	defer ch.Close()

	err := ch.ExchangeDeclare(
		"logs_topic", // name
		"topic", // type
		true, // durable
		false, // auto-deleted
		false, // internal
		false, // no-wait
		nil, // arguments
	)
	utils.FailOnError(err, "Failed to declare an exchange")

	body := utils.BodyFrom(os.Args)

	err = ch.Publish(
		"logs_topic", // exchange
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
