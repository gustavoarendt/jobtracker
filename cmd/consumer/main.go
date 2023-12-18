package main

import (
	"github.com/gustavoarendt/jobtracker/pkg/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	ch, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	output := make(chan amqp.Delivery)

	go rabbitmq.Consume(ch, output, "oportunities")

	for message := range output {
		println(string(message.Body))
		message.Ack(false)
	}
}
