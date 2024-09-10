package rabbit

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

type Queue struct {
	Connect   *amqp.Connection
	QueueName string
	Callback  func([]byte)
}

// if err output
func failOnError(err error, msg string) {
	if err != nil {
		log.Printf("%s: %s", msg, err)
	}
}

func (t *Queue) Consumer() {

	channel, err := t.Connect.Channel()
	failOnError(err, "Failed to open a channel")
	defer channel.Close()

	//declare a queue
	q, err := channel.QueueDeclare(
		t.QueueName, // queue name
		false,       // durable
		false,       // delete when unused
		false,       // exclusive:when a consumer close connection，delete the queue
		false,       // no-wait
		nil,         // arguments
	)
	failOnError(err, "Failed to declare a queue")

	//consume message
	msgs, err := channel.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	//create a go channel
	forever := make(chan bool)

	//run by gorountine
	go func() {
		for d := range msgs {
			t.Callback(d.Body) //call the callback
		}
	}()

	//waiting the mesage forever
	<-forever
}
