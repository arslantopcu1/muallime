package rabbit

import amqp "github.com/rabbitmq/amqp091-go"

type Topic struct {
	Connect   *amqp.Connection
	QueueName string
	TopicName string
	Callback  func([]byte)
}

func (t *Topic) Consumer() {

	ch, err := t.Connect.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		t.TopicName, // name
		"topic",     // type
		true,        // durable
		false,       // auto-deleted
		false,       // internal
		false,       // no-wait
		nil,         // arguments
	)
	failOnError(err, "Failed to declare an exchange")

	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when usused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")

	err = ch.QueueBind(
		q.Name,      // queue name
		t.QueueName, // routing key
		t.TopicName, // exchange
		false,
		nil)

	failOnError(err, "Failed to bind a queue")

	//consume message
	msgs, err := ch.Consume(
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
