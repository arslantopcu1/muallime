package rabbit

import (
	"errors"
	amqp "github.com/rabbitmq/amqp091-go"
	"swagger/database"
)

type PublisherQueue struct {
	QueueName string
	Message   []byte
}

func (t *PublisherQueue) Publish() error {

	conn := database.MQ

	if conn == nil {
		return errors.New("Rabbitmq")
	}

	ch, err := conn.Channel()
	defer ch.Close()

	if err != nil {
		return errors.New("Failed to open a channel")
	}

	q, err := ch.QueueDeclare(
		t.QueueName, // name
		false,       // durable
		false,       // delete when unused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)

	if err != nil {
		return errors.New("Failed to declare a queue")
	}

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        t.Message,
		})

	if err != nil {
		return errors.New("Failed to publish a message")
	}

	return nil

}
