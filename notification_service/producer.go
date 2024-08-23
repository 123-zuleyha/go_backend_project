// producer.go
package rabbitmq

import "github.com/streadway/amqp"

func PublishMessage(ch *amqp.Channel, queueName string, body string) {
	q, err := ch.QueueDeclare(
		queueName,
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	FailOnError(err, "Failed to declare a queue")

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	FailOnError(err, "Failed to publish a message")
}
