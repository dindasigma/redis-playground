package main

// producer: POST http://localhost:4151/pub?topic=clicks

import (
	"errors"
	"github.com/nsqio/go-nsq"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type NoopNSQLogger struct{}

func (l *NoopNSQLogger) Output(calldepth int, s string) error {
	return nil
}

type MessageHandler struct{}

func (h *MessageHandler) HandleMessage(m *nsq.Message) error {
	if len(m.Body) == 0 {
		return errors.New("body is blank re-enqueue message")
	}

	log.Print(m.Body)
	return nil
}

func main() {
	config := nsq.NewConfig()

	consumer, err := nsq.NewConsumer("clicks", "metrics", config)
	if err != nil {
		log.Fatal(err)
	}

	consumer.ChangeMaxInFlight(200)

	consumer.SetLogger(
		&NoopNSQLogger{},
		nsq.LogLevelError,
	)

	consumer.AddConcurrentHandlers(
		&MessageHandler{},
		20,
	)

	if err := consumer.ConnectToNSQD("localhost:4150"); err != nil {
		log.Fatal(err)
	}

	/*command := nsq.Command{
		[]byte("test"),
		[]byte("test"),
		[]byte("test"),
	}*/


	shutdown := make(chan os.Signal, 2)
	signal.Notify(shutdown, syscall.SIGINT)

	for {
		select {
		case <-consumer.StopChan:
			return
		case <-shutdown:
			consumer.Stop()
		}
	}

}
