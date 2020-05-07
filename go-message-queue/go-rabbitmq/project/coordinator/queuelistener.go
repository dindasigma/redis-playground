package coordinator

// coordinator responsibility:
// 1. Interact with the sensor via message broker
// 2.

import (
	"bytes"
	"encoding/gob"
	"fmt"

	"github.com/dindasigma/go-rabbitmq/project/dto"
	"github.com/dindasigma/go-rabbitmq/project/qutils"
	"github.com/streadway/amqp"
)

const url = "amqp://guest:guest@localhost:5672"

type QueueListener struct {
	conn *amqp.Connection
	ch *amqp.Channel
	sources map[string]<-chan amqp.Delivery
}

func NewQueueListener() *QueueListener {
	ql := QueueListener{
		sources: make(map[string]<-chan amqp.Delivery),
	}

	ql.conn, ql.ch = qutils.GetChannel(url)

	return &ql
}

func (ql *QueueListener) ListenForNewSource() {
	// create queue with no name, rabbit-mq will generate the name automatically
	q := qutils.GetQueue("", ql.ch)

	// bind the fan-out queue
	ql.ch.QueueBind(
		q.Name,
		"",
		"amq.fanout",
		false,
		nil)

	msgs, _ := ql.ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
		)

	for msg := range msgs {
		sourceChan, _ := ql.ch.Consume(
			string(msg.Body),
			"",
			true,
			false,
			false,
			false,
			nil)

		if ql.sources[string(msg.Body)] == nil {
			ql.sources[string(msg.Body)] = sourceChan

			go ql.AddListener(sourceChan)
		}
	}
}

func (ql *QueueListener) AddListener(msgs <-chan amqp.Delivery) {
	for msg := range msgs {
		r := bytes.NewReader(msg.Body)
		d := gob.NewDecoder(r)
		sd := new(dto.SensorMessage)
		d.Decode(sd)

		fmt.Printf("Received message: %v\n", sd)
	}
}