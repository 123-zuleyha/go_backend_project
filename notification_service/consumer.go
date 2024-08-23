package rabbitmq

import (
	"log"

	"github.com/streadway/amqp"
)

func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func ConsumeMessages(ch *amqp.Channel, queueName string) {
	msgs, err := ch.Consume(
		queueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	FailOnError(err, "Failed to register a consumer")

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()
}
