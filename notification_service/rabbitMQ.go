// rabbitmq.go
package rabbitmq

import (
	"log"

	"github.com/streadway/amqp"
)

func FailOnErrorr(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func Connect() (*amqp.Connection, *amqp.Channel) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	FailOnError(err, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()
	FailOnError(err, "Failed to open a channel")

	return conn, ch
}
